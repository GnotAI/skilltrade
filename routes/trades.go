package routes

import (
	"github.com/GnotAI/skilltrade/db"
	"github.com/GnotAI/skilltrade/internal/trades"
	"github.com/GnotAI/skilltrade/middleware"
	"github.com/go-chi/chi/v5"
)

var tradeRepo = trades.NewTradeRepository(db.DB)
var tradeService = trades.NewTradeService(tradeRepo)
var tradeHandler = trades.NewTradeHandler(tradeService)

func tradeRoutes() *chi.Mux {
  r := chi.NewRouter()

  r.Group(func (r chi.Router)  {
    r.Use(middleware.AuthMiddleware)
    
    r.Post("/", tradeHandler.CreateTrade)
  })

  return r
}
