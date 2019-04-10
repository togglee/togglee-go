package pkg

import (
	"github.com/feaggle/feaggle/pkg/drivers"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)


func Test_Integration_Toggle(t *testing.T) {

	os.Setenv("FAKE_TOGGLE_ENVIRONMENT_VARIABLE", "true")
	feaggle := FeaggleBuilder{}.AddDriver(&drivers.ReleaseEnvironmentDriver{}).Create()
	assert.True(t, feaggle.IsActive("FAKE_TOGGLE_ENVIRONMENT_VARIABLE"), "if is true, returns true")
}