package sci_operations

import "github.com/goarchitecture/go-simple-calculator/calculator/operations"

func All() []operations.Operation {
	return []operations.Operation{
		NewSin(),
	}
}
