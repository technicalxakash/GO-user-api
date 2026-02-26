package config

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret      string // Store this in environment variable in production
	ExpiryHours int
}

var JWT = JWTConfig{
	Secret:      "your-super-secret-jwt-key-change-in-production", // CHANGE THIS IN PRODUCTION
	ExpiryHours: 24,
}
