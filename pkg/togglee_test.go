package pkg_test

import (
	"github.com/bxcodec/faker/v3"
	. "github.com/togglee/togglee-go/pkg"
	"github.com/togglee/togglee-go/pkg/models/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func boolToPointer(b bool) *bool {
	return &b
}

func Test_Toggle_Is_False_If_Not_Exist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name, make(map[string]interface{})).Return(nil)

	togglee := ToggleeBuilder{}.AddDriver(driver).Create()

	assert.False(t, togglee.IsActive(name, make(map[string]interface{})), "defaults to false if does not exist")
}

func Test_Toggle_Is_True_If_Exist_In_Driver(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name, make(map[string]interface{})).Return(boolToPointer(true))

	togglee := ToggleeBuilder{}.AddDriver(driver).Create()

	assert.True(t, togglee.IsActive(name, make(map[string]interface{})), "is true in driver")
}

func Test_Toggle_Is_False_If_Exist_In_Driver(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name, make(map[string]interface{})).Return(boolToPointer(false))

	togglee := ToggleeBuilder{}.AddDriver(driver).Create()

	assert.False(t, togglee.IsActive(name, make(map[string]interface{})), "is false in driver")
}

func Test_Toggle_Multiple_Drivers_False(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	otherDriver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name, make(map[string]interface{})).Return(nil)
	otherDriver.EXPECT().IsActive(name, make(map[string]interface{})).Return(boolToPointer(true))

	togglee := ToggleeBuilder{}.AddDriver(driver).AddDriver(otherDriver).Create()

	assert.True(t, togglee.IsActive(name, make(map[string]interface{})), "is false in driver")
}

func Test_Toggle_Multiple_Drivers_False_Not_Exist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	otherDriver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name, make(map[string]interface{})).Return(nil)
	otherDriver.EXPECT().IsActive(name, make(map[string]interface{})).Return(nil)

	togglee := ToggleeBuilder{}.AddDriver(driver).AddDriver(otherDriver).Create()

	assert.False(t, togglee.IsActive(name, make(map[string]interface{})), "is false in driver")
}

func Test_Toggle_Multiple_Drivers_False_Exist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	otherDriver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name, make(map[string]interface{})).Return(boolToPointer(true))
	otherDriver.EXPECT().IsActive(name, make(map[string]interface{})).Return(boolToPointer(false))

	togglee := ToggleeBuilder{}.AddDriver(driver).AddDriver(otherDriver).Create()

	assert.False(t, togglee.IsActive(name, make(map[string]interface{})), "is false in driver")
}

func Test_Toggle_Multiple_Drivers_False_Exist_Does_Not_Call_Next(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	otherDriver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name, make(map[string]interface{})).Return(boolToPointer(false))
	otherDriver.EXPECT().IsActive(name, make(map[string]interface{})).Return(boolToPointer(true)).Times(0)

	togglee := ToggleeBuilder{}.AddDriver(driver).AddDriver(otherDriver).Create()

	assert.False(t, togglee.IsActive(name, make(map[string]interface{})), "is false in driver")
}