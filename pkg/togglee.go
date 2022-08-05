package pkg

import "github.com/togglee/togglee-go/pkg/models"

// ToggleeBuilder -- toggle resolver builder
type ToggleeBuilder struct {
	drivers []models.Driver
}

// Togglee -- toggle resolver
type Togglee struct {
	drivers []models.Driver
}

// AddDriver -- Adds a driver to the unes used by the toggle resolver
func (h ToggleeBuilder) AddDriver(driver models.Driver) ToggleeBuilder {
	h.drivers = append(h.drivers, driver)
	return h
}

// Create -- Creates an instance of the feature toggle resolver
func (h ToggleeBuilder) Create() *Togglee {
	return &Togglee {
		drivers: h.drivers,
	}
}

// IsActive -- retrieves the value of a toggle from the drivers
func (h *Togglee) IsActive(name string, context map[string]interface{}) bool {
	result := false
	for _, driver := range h.drivers {
		resultLocal := driver.IsActive(name, context)
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