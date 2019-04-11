package drivers

import (
	"github.com/feaggle/feaggle/pkg/models"
)

// ReleaseDriver -- release toggles driver resolver
type ReleaseRemoteDriver struct {
	Cache models.ToggleCache
}

// IsActive -- release toggles driver resolver function
func (h *ReleaseRemoteDriver) IsActive(name string) *bool {
	toggles := h.Cache.Toggles()
	for _, value := range toggles.Releases {
		if value.Name == name {
			return &value.Active
		}
	}
	return nil
}
