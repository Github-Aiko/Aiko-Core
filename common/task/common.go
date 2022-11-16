package task

import "github.com/Github-Aiko/Aiko-Core/common"

// Close returns a func() that closes v.
func Close(v interface{}) func() error {
	return func() error {
		return common.Close(v)
	}
}
