package server

type Server interface {
	Run()
	ListenError() <-chan error
}
