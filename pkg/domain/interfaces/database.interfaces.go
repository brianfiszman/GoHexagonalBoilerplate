package interfaces

type Database interface {
	ConnectDatabase()
	Ping() error
}
