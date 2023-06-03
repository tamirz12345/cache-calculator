package storage

// CalculatorStorage interface for storage for the calculator
type CalculatorStorage interface {
	Get(key string) (float64, bool)
	Set(key string, val float64)
	GetSize() int
}
