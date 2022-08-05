//go:generate mockgen -source=driver.go -package mocks -destination  mocks/driver.go

package models

// Driver -- implementation of single resolver for toggles
type Driver interface {
	IsActive(string, map[string]interface{}) *bool
}
