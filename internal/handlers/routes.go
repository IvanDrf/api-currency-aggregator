package handlers

func (s *Server) RegisterRoutes() {
	s.server.POST("/", PostHandler)
	s.server.GET("/", GetHandler)

}
