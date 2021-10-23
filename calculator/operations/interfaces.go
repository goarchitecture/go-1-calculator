package operations

type Operation interface {
	Setup(operands ...float64) Operation

	Match(userOperation string) bool
	Do() (result float64, err error)
}
