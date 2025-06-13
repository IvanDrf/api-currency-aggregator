package handlers

func (s *Server) RegisterRoutes() {
	s.Server.POST("/", PostHandler)
	s.Server.GET("/", GetHandler)

}
