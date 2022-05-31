package healthCheck

type Service interface {
	Ping() map[string]string
}
