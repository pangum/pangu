package rest

import (
	`fmt`
	`net/http`

	`github.com/pangum/pangu/app`
)

// Server RESTFul服务器
// 这里只做演示用，实际开发中请使用https://github.com/pangum/web开发
// 也可以使用纯属Echo或者Gin以及其它框架来做开发
type Server struct {
	app.NamedServe
}

func newServer() *Server {
	return &Server{
		NamedServe: app.NewNamedServe(`RESTFul服务器`),
	}
}

func (s *Server) Start() error {
	http.HandleFunc(`/`, s.indexHandler)

	return http.ListenAndServe(`:8000`, nil)
}

func (s *Server) indexHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, `hello world`)
}
