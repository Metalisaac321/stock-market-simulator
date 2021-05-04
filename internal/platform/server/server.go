package server

import (
	"fmt"
	"log"

	"github.com/Metalisaac321/stock-market-simulator/internal/creating"
	"github.com/Metalisaac321/stock-market-simulator/internal/platform/server/handler/accounts"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	accountService creating.AccountService
}

func New(host string, port uint, accountService creating.AccountService) Server {
	srv := Server{
		engine:         gin.New(),
		httpAddr:       fmt.Sprintf("%s:%d", host, port),
		accountService: accountService,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.POST("/accounts", accounts.CreateHandler(s.accountService))
}
