package core

type Calculator interface {
	Calculate(string) (string, error)
}
