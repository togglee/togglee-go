package pkg

import (
	"fmt"
	"testing"
)


func Test_Toggle(t *testing.T) {
	feaggle := FeaggleBuilder{}.Create()
	fmt.Println("Value:", feaggle.IsActive("pepe"))
}