package drivers

import (
	"os"
	"strconv"
)

// ReleaseDriver -- release toggles driver resolver
type ReleaseEnvironmentDriver struct{ }

// IsActive -- release toggles driver resolver function
func (h *ReleaseEnvironmentDriver) IsActive(name string, context map[string]interface{}) *bool {
	result, error := strconv.ParseBool(os.Getenv(name))
	if error != nil {
		return nil
	}
	return  &result
}
