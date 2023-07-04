package reporter

type result int

const (
	RequestNotSent result = iota
	RequestSent
	RequestError
)

func (r result) String() string {
	return [...]string{"Not Sent", "Sent", "Error"}[r]
}
