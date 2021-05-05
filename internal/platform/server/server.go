package server

import (
	"fmt"
	"log"

	"github.com/Metalisaac321/stock-market-simulator/internal/account/application"
	"github.com/Metalisaac321/stock-market-simulator/internal/platform/server/handler/accounts"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	createAccount     application.CreateAccount
	searchAllAccounts application.SearchAllAccounts
}

func New(host string, port uint, createAccount application.CreateAccount, searchAllAccounts application.SearchAllAccounts) Server {
	srv := Server{
		engine:            gin.New(),
		httpAddr:          fmt.Sprintf("%s:%d", host, port),
		createAccount:     createAccount,
		searchAllAccounts: searchAllAccounts,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.POST("/accounts", accounts.CreateHandler(s.createAccount))
	s.engine.GET("/accounts", accounts.SearchAllHandler(s.searchAllAccounts))
}
