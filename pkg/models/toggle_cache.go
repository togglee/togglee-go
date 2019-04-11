////go:generate mockgen -self_package github.com/feaggle/feaggle/pkg/models -source toggle_cache.go -package mocks -destination mocks/toggle_cache.go

package models

type ToggleCache interface {
	Toggles() *Toggles
}