package types

import "context"

type (
	CacheService interface {
		// Set method sets data to cache
		Set(ctx context.Context, key string, value []byte) error

		// Get method gets data from cache
		Get(ctx context.Context, key string) ([]byte, error)

		// SetUser sets user data to cache
		SetUser(ctx context.Context, user *User) error

		// GetUser gets user data from cache
		GetUser(ctx context.Context, name string, rollNo int64) (*User, error)
	}

	// User is a sample user struct
	User struct {
		Name     string `json:"name"`
		Class    string `json:"class"`
		RollNum  int64  `json:"roll_num"`
		Metadata []byte `json:"metadata"`
	}
)
