package controllers

func (s *Server) initializeRoutes() {
	v1 := s.Router.Group("/api/v1")
	{
		// Main route to ask for the beginning and ending of a flight path
		v1.POST("/flightPath", s.GetStartAndDestination)
	}
}
