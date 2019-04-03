package models

// Driver -- implementation of single resolver for toggles
type Driver interface {
	IsActive(string) *bool
}
