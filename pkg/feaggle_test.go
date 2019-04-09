package pkg_test

import (
	"github.com/bxcodec/faker/v3"
	. "github.com/feaggle/feaggle/pkg"
	"github.com/feaggle/feaggle/pkg/models/mocks"
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
	driver.EXPECT().IsActive(name).Return(nil)

	feaggle := FeaggleBuilder{}.AddDriver(driver).Create()

	assert.False(t, feaggle.IsActive(name), "defaults to false if does not exist")
}

func Test_Toggle_Is_True_If_Exist_In_Driver(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name).Return(boolToPointer(true))

	feaggle := FeaggleBuilder{}.AddDriver(driver).Create()

	assert.True(t, feaggle.IsActive(name), "is true in driver")
}

func Test_Toggle_Is_False_If_Exist_In_Driver(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name).Return(boolToPointer(false))

	feaggle := FeaggleBuilder{}.AddDriver(driver).Create()

	assert.False(t, feaggle.IsActive(name), "is false in driver")
}

func Test_Toggle_Multiple_Drivers_False(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	otherDriver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name).Return(nil)
	otherDriver.EXPECT().IsActive(name).Return(boolToPointer(true))

	feaggle := FeaggleBuilder{}.AddDriver(driver).AddDriver(otherDriver).Create()

	assert.True(t, feaggle.IsActive(name), "is false in driver")
}

func Test_Toggle_Multiple_Drivers_False_Not_Exist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	otherDriver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name).Return(nil)
	otherDriver.EXPECT().IsActive(name).Return(nil)

	feaggle := FeaggleBuilder{}.AddDriver(driver).AddDriver(otherDriver).Create()

	assert.False(t, feaggle.IsActive(name), "is false in driver")
}

func Test_Toggle_Multiple_Drivers_False_Exist(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	otherDriver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name).Return(boolToPointer(true))
	otherDriver.EXPECT().IsActive(name).Return(boolToPointer(false))

	feaggle := FeaggleBuilder{}.AddDriver(driver).AddDriver(otherDriver).Create()

	assert.False(t, feaggle.IsActive(name), "is false in driver")
}

func Test_Toggle_Multiple_Drivers_False_Exist_Does_Not_Call_Next(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	name := faker.UUIDHyphenated()
	driver := mocks.NewMockDriver(controller)
	otherDriver := mocks.NewMockDriver(controller)
	driver.EXPECT().IsActive(name).Return(boolToPointer(false))
	otherDriver.EXPECT().IsActive(name).Return(boolToPointer(true)).Times(0)

	feaggle := FeaggleBuilder{}.AddDriver(driver).AddDriver(otherDriver).Create()

	assert.False(t, feaggle.IsActive(name), "is false in driver")
}