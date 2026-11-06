package server

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Start() {}

func (server *Server) Stop() {}

func (server *Server) openWebSocket() {}

func (server *Server) closeWebSocket() {}
