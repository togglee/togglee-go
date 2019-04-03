package pkg

import "github.com/feaggle/feaggle/pkg/models"

// FeaggleBuilder -- toggle resolver builder
type FeaggleBuilder struct {
	drivers []models.Driver
}

// Feaggle -- toggle resolver
type Feaggle struct {
	drivers []models.Driver
}

// AddDriver -- Adds a driver to the unes used by the toggle resolver
func (h FeaggleBuilder) AddDriver(driver models.Driver) FeaggleBuilder {
	h.drivers = append(h.drivers, driver)
	return h
}

// Create -- Creates an instance of the feature toggle resolver
func (h FeaggleBuilder) Create() Feaggle {
	return Feaggle {
		drivers: h.drivers,
	}
}

// IsActive -- retrieves the value of a toggle from the drivers
func (h *Feaggle) IsActive(name string) bool {
	result := false
	for _, driver := range h.drivers {
		resultLocal := driver.IsActive(name)
		if resultLocal == nil {
			continue
		}
		if !*resultLocal {
			return false
		}
		result = true
	}
	return result
}