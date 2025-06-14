package handlers

func (s *Server) RegisterRoutes() {
	s.Server.GET("/sources", GetListHandler)
	s.Server.GET("/currency", GetCurrencyHandler)

}
