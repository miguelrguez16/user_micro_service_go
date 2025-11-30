package config

import (
	"context"
	"os"
	"testing"
)

// Test_ConnectDB tests ConnectDB with various configurations
func Test_ConnectDB(t *testing.T) {
	tests := []struct {
		name               string
		env                string
		mongoURI           string
		dbName             string
		expectedProduction bool
	}{
		{
			name:               "production mode",
			env:                "production",
			mongoURI:           "mongodb://localhost:27017",
			dbName:             "test_users_db",
			expectedProduction: true,
		},
		{
			name:               "development mode",
			env:                "development",
			mongoURI:           "mongodb://localhost:27017",
			dbName:             "test_users_db",
			expectedProduction: false,
		},
		{
			name:               "default mode when ENV not set",
			env:                "",
			mongoURI:           "mongodb://localhost:27017",
			dbName:             "test_users_db",
			expectedProduction: false,
		},
		{
			name:               "empty DB_NAME",
			env:                "",
			mongoURI:           "mongodb://localhost:27017",
			dbName:             "",
			expectedProduction: false,
		},
		{
			name:               "custom MONGO_URI",
			env:                "production",
			mongoURI:           "mongodb://test-host:27017",
			dbName:             "custom_db",
			expectedProduction: true,
		},
		{
			name:               "invalid MONGO_URI",
			env:                "",
			mongoURI:           "invalid://uri",
			dbName:             "test_db",
			expectedProduction: false,
		},
		{
			name:               "missing MONGO_URI",
			env:                "",
			mongoURI:           "",
			dbName:             "test_db",
			expectedProduction: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup environment
			if tt.env != "" {
				os.Setenv("ENV", tt.env)
			} else {
				os.Unsetenv("ENV")
			}

			if tt.mongoURI != "" {
				os.Setenv("MONGO_URI", tt.mongoURI)
			} else {
				os.Unsetenv("MONGO_URI")
			}

			if tt.dbName != "" {
				os.Setenv("DB_NAME", tt.dbName)
			} else {
				os.Unsetenv("DB_NAME")
			}

			// Execute
			db, err, isProduction := ConnectDB()

			// Verify production flag
			if isProduction != tt.expectedProduction {
				t.Errorf("Expected isProduction=%v, got %v", tt.expectedProduction, isProduction)
			}

			// Log any connection error (expected in test environment)
			if err != nil {
				t.Logf("Connection error (expected in test environment): %v", err)
			}

			// Cleanup
			if db != nil && db.Client() != nil {
				db.Client().Disconnect(context.TODO())
			}
		})
	}
}
