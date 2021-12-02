package util

// ConcName  identifica a las peticiones concurrentes
type ConcName string

// const (
// 	concFindAccountByEmail    ConcName = "find-account-by-email"
// 	concFindAccountByUsername ConcName = "find-account-by-username"
// )

// ConcurrencyData
type ConcurrencyData struct {
	Err  error
	Name ConcName
}
