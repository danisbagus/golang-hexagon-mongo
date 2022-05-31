package healthCheck

type Repository interface {
	Ping() error
}
