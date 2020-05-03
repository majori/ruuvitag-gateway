package ruuvi

type DatabaseAdapter interface {
	Save(*Measurement) error
}
