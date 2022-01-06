package util

// ConcName  identifica a las peticiones concurrentes
type ConcName string

// ConcurrencyData
type ConcurrencyData struct {
	Err  error
	Name ConcName
}
