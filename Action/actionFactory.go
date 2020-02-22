package Action

import (
	"fmt"
)

const (
	Initialize = "initialize_action"
)

func GetAction(name string) actionInterface {
	switch name {
	case Initialize:
		return new(InitializeAction)
	default:
		panic(fmt.Sprintf("Action %d not recognized\n", name))
	}
}