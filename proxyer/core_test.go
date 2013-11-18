package proxyer

import (
	"testing"
)

func Test_Core_Std(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	t.Error("this is just hardcoded as an error.") //Indicate that this test failed and log the string as info
}
