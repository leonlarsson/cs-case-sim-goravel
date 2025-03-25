package factories

import (
	"math/rand/v2"
)

type UnboxFactory struct {
}

// Definition Define the model's default state.
func (f *UnboxFactory) Definition() map[string]any {
	return map[string]any{
		"CaseId":     "crate-4904",
		"ItemId":     "skin-461956",
		"IsStatTrak": rand.IntN(100) < 10,
		"UnboxerId":  "totally a UUID bro",
	}
}
