package handlers

func (s *Server) RegisterRoutes() {
	s.Server.GET("/", GetListHandler)
	s.Server.GET("/currency", GetCurrencyHandler)

}
