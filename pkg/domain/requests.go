package domain

type Requests interface {
	Get(string) (string, error)
}
