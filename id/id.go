/*
	Package id provides unique ID generation.
*/
package id

import(
	"github.com/satori/go.uuid"
)

// Returns a UUID encoded as a string.
func NewID() (string, error) {
	return uuid.NewV4().String(), nil
}
