package http

//DAO for http methods
type DAO interface {
	Get(url string) ([]byte, int, error)
}
