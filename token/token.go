package token

import "time"

// Token is either access token or refresh token, you can set/read value, check is expired.
//
type Token interface {

	// ToString return string with expired time, after expired time the token will not read from string
	//
	//	expired := time.Now().UTC().Add(60 * time.Second)
	//	str := token.ToString(expired)
	//
	ToString(expired time.Time) (string, error)

	// Get return value from key
	//
	//	value := token.Get("UserID")
	//
	Get(key string) string

	// Set return value to key
	//
	//	token.Set("UserID","aa")
	//
	Set(key, value string)

	// Delete key
	//
	//	token.Delete("UserID")
	//
	Delete(key string)
}
