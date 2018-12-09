package mtest

import (
	"fmt"
)

// Hi - return 'Hi' message
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}
