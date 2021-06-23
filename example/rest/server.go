package rest

import (
	`fmt`
	`net/http`
)

// Server RESTFul服务器
// 这里只做演示用，实际开发中请使用https://github.com/storezhang/echox开发
// 也可以使用纯属echo或者gin以及其它框架来做开发
type Server struct{}

func newServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	http.HandleFunc("/", s.indexHandler)

	return http.ListenAndServe(":8000", nil)
}

func (s *Server) indexHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "hello world")
}

func (s *Server) Name() string {
	return "RESTFul服务器"
}
