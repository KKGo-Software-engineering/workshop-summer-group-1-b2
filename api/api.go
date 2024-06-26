package api

import (
	"database/sql"

	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/KKGo-Software-engineering/workshop-summer/api/eslip"
	"github.com/KKGo-Software-engineering/workshop-summer/api/health"
	"github.com/KKGo-Software-engineering/workshop-summer/api/mlog"
	"github.com/KKGo-Software-engineering/workshop-summer/api/spender"
	"github.com/KKGo-Software-engineering/workshop-summer/api/transaction"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	*echo.Echo
}

func New(db *sql.DB, cfg config.Config, logger *zap.Logger) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(mlog.Middleware(logger))

	v1 := e.Group("/api/v1")

	v1.GET("/slow", health.Slow)
	v1.GET("/health", health.Check(db))
	v1.POST("/upload", eslip.Upload)

	// v1.Use(middleware.BasicAuth(AuthCheck))

	{
		h := spender.New(cfg.FeatureFlag, db)
		v1.GET("/spenders", h.GetAll)
		v1.POST("/spenders", h.Create)
	}

	{
		h := transaction.New(cfg.FeatureFlag, db)
		v1.GET("/expenses/summary", h.GetSummaryExpensesHandler)
		v1.GET("/incomes/summary", h.GetSummaryExpensesHandler)
		v1.GET("/balance", h.GetSummaryBalanceHandler)
	}
	{
		h := transaction.New(cfg.FeatureFlag, db)
		v1.GET("/transactions", h.GetAllTransaction)
		v1.GET("/transactions?page=1&limit=1&transaction_type=expense", h.GetFilteredTransaction)
		//v1.POST("/spenders", h.Create)
		h = transaction.New(cfg.FeatureFlag, db)
		v1.POST("/transactions", h.CreateTransaction)
		v1.PUT("/transactions/:id", h.UpdateTransaction)
	}
	return &Server{e}
}
