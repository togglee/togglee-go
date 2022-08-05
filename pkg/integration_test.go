package pkg

import (
	"github.com/togglee/togglee-go/pkg/drivers"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)


func Test_Integration_Toggle(t *testing.T) {

	os.Setenv("FAKE_TOGGLE_ENVIRONMENT_VARIABLE", "true")
	togglee := ToggleeBuilder{}.AddDriver(&drivers.ReleaseEnvironmentDriver{}).Create()
	assert.True(t, togglee.IsActive("FAKE_TOGGLE_ENVIRONMENT_VARIABLE", make(map[string]interface{})), "if is true, returns true")
}