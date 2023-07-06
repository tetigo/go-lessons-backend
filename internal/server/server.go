package server

import (
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/tetigo/go-lessons-backend/internal/category"
	"github.com/tetigo/go-lessons-backend/internal/product"
)

type Server struct {
	engine *gin.Engine
	port uint
}

type Config struct {
	Port uint
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	server := &Server{
		engine: engine,
		port: config.Port,
	}
	engine.GET("/ping", server.HealthCheck)
	engine.GET("/categories", server.Categories)
	engine.GET("/products", server.Products)
	return server, nil
}

func (s *Server) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Server) Run() error{
	return s.engine.Run(fmt.Sprintf(":%d", s.port))
}


func (s *Server) Categories(c *gin.Context) {
	categories := []category.Category{
		{
			ID:          "12",
			Name:        "Category1",
			Description: "Category1 Desc",
		},
		{
			ID:          "22",
			Name:        "Category2",
			Description: "Category2 Desc",
		},
	}

	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.JSON(200, categories)
}

func (s *Server) Products(c *gin.Context) {
	products := []product.Product{
		{
			ID:               "123",
			Name:             "Product1",
			Description:      "Product1 desc",
			PriceVATExcluded: money.New(100, "USD"),
			VAT:              money.New(200, "USD"),
		},
		{
			ID:               "124",
			Name:             "Product2",
			Description:      "Product2 desc",
			PriceVATExcluded: money.New(200, "USD"),
			VAT:              money.New(300, "USD"),
		},
		{
			ID:               "125",
			Name:             "Product3",
			Description:      "Product3 desc",
			PriceVATExcluded: money.New(500, "USD"),
			VAT:              money.New(600, "USD"),
		},
	}
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.JSON(200, products)
}
