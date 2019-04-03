package pkg

import (
	"fmt"
	"github.com/feaggle/feaggle/pkg/drivers"
	"testing"
)


func Test_Toggle(t *testing.T) {
	feaggle := FeaggleBuilder{}.AddDriver(&drivers.ReleaseDriver{}).Create()
	fmt.Println("Value:", feaggle.IsActive("pepe"))
}