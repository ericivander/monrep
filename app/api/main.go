package main

import (
	"net/http"

	"github.com/ericivander/monrep/config"
	"github.com/ericivander/monrep/flow"
	"github.com/ericivander/monrep/handler"
	"github.com/ericivander/monrep/repository/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.NewConfig()

	postgresDb, err := postgres.NewPostgres(&cfg.Postgres)
	if err != nil {
		panic(err)
	}

	categoryRepo := postgres.NewCategory(postgresDb.Db)
	incomeRepo := postgres.NewIncome(postgresDb.Db)
	expenseRepo := postgres.NewExpense(postgresDb.Db)

	reportFlow := flow.NewReportProvider(categoryRepo, incomeRepo, expenseRepo)

	reportHandler := handler.NewReportHandler(reportFlow)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/reports", reportHandler.GetReports)

	// Start server
	e.Logger.Fatal(e.Start(":2234"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
