package drivers_test

import (
	"github.com/feaggle/feaggle/pkg/drivers"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Release_Environment_Driver(t *testing.T) {
	t.Run("takes true value from environment", func(t *testing.T) {

		os.Setenv("FAKE_TOGGLE_VARIABLE", "true")
		driver := drivers.ReleaseEnvironmentDriver{}
		result := driver.IsActive("FAKE_TOGGLE_VARIABLE")
		assert.True(t, *result, "true if exist as bool")
	})

	t.Run("takes false value from environment", func(t *testing.T) {

		os.Setenv("FAKE_TOGGLE_VARIABLE", "false")
		driver := drivers.ReleaseEnvironmentDriver{}
		result := driver.IsActive("FAKE_TOGGLE_VARIABLE")
		assert.False(t, *result, "false if exist as bool")
	})

	t.Run("ignored non boolean values", func(t *testing.T) {

		os.Setenv("FAKE_TOGGLE_VARIABLE", "SomeOtherStuff")
		driver := drivers.ReleaseEnvironmentDriver{}
		result := driver.IsActive("FAKE_TOGGLE_VARIABLE")
		assert.Nil(t, result, "nil if not a bool")
	})

	t.Run("ignored non existing environments", func(t *testing.T) {
		driver := drivers.ReleaseEnvironmentDriver{}
		result := driver.IsActive("OTHER_FAKE_TOGGLE_VARIABLE")
		assert.Nil(t, result, "nil if does not exist")
	})

}
