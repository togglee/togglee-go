////go:generate mockgen -self_package github.com/togglee/togglee-go/pkg/models -source toggle_cache.go -package mocks -destination mocks/toggle_cache.go

package models

type ToggleCache interface {
	Toggles() *Toggles
}