package idgen

import "github.com/jaevor/go-nanoid"

var (
	generateAlphaNum20U func() string
)

func init() {
	generateAlphaNum20U, _ = nanoid.Custom("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 20)
}

// AlphaNum20U alpha numeric 20 characters all uppercase
func AlphaNum20U() string {
	return generateAlphaNum20U()
}
