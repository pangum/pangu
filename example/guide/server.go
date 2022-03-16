package guide

type server struct {
	constructor *constructor
}

func newServer(constructor *constructor) *server {
	return &server{
		constructor: constructor,
	}
}
