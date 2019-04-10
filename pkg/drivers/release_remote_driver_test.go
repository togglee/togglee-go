package drivers_test

import (
	"github.com/feaggle/feaggle/pkg/drivers"
	. "github.com/feaggle/feaggle/pkg/models"
	"github.com/feaggle/feaggle/pkg/models/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Release_Remote_Driver(t *testing.T) {
	t.Run("takes true value from cache", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		toggleCacheMock := mocks.NewMockToggleCache(controller)
		toggleCacheMock.EXPECT().Validate().Times(1)
		toggles := &Toggles{
			Releases:
			[]ReleaseToggle{
				{Name: "SOME_OTHER_TOGGLE_NAME", Active: false},
				{Name: "TOGGLE_NAME", Active: true},
			},
		}
		toggleCacheMock.EXPECT().Cache().Return(toggles).Times(1)
		driver:= drivers.ReleaseRemoteDriver{
			Cache: toggleCacheMock,
		}
		assert.True(t, *driver.IsActive("TOGGLE_NAME"))
	})

	t.Run("takes false value from cache", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		toggleCacheMock := mocks.NewMockToggleCache(controller)
		toggleCacheMock.EXPECT().Validate().Times(1)
		toggles := &Toggles{
			Releases:
			[]ReleaseToggle{
				{Name: "SOME_OTHER_TOGGLE_NAME", Active: true},
				{Name: "TOGGLE_NAME", Active: false},
			},
		}
		toggleCacheMock.EXPECT().Cache().Return(toggles).Times(1)
		driver:= drivers.ReleaseRemoteDriver{
			Cache: toggleCacheMock,
		}
		assert.False(t, *driver.IsActive("TOGGLE_NAME"))
	})

	t.Run("returns false if not found", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		toggleCacheMock := mocks.NewMockToggleCache(controller)
		toggleCacheMock.EXPECT().Validate().Times(1)
		toggles := &Toggles{
			Releases: []ReleaseToggle{
				{Name: "SOME_OTHER_TOGGLE_NAME", Active: true},
				{Name: "SOME_MORE_TOGGLE_NAME", Active: true},
			},
		}
		toggleCacheMock.EXPECT().Cache().Return(toggles).Times(1)
		driver:= drivers.ReleaseRemoteDriver{
			Cache: toggleCacheMock,
		}
		assert.Nil(t, driver.IsActive("TOGGLE_NAME"))
	})

	t.Run("returns false if not valid toggles", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		toggleCacheMock := mocks.NewMockToggleCache(controller)
		toggleCacheMock.EXPECT().Validate().Times(1)
		toggles := &Toggles{
			Releases: nil,
		}
		toggleCacheMock.EXPECT().Cache().Return(toggles).Times(1)
		driver:= drivers.ReleaseRemoteDriver{
			Cache: toggleCacheMock,
		}
		assert.Nil(t, driver.IsActive("TOGGLE_NAME"))
	})
}
