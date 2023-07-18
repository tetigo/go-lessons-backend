package server

import (
	"fmt"
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/tetigo/go-lessons-backend/internal/category"
	"github.com/tetigo/go-lessons-backend/internal/product"
)

type Server struct {
	Engine *gin.Engine
	port   uint
	allowedOrigin string
	model string
}

type Config struct {
	Port uint
	AllowedOrigin string
}

func New(config Config) (*Server, error) {
	engine := gin.Default()
	server := &Server{
		Engine: engine,
		port:   config.Port,
		allowedOrigin: config.AllowedOrigin,
		model: "Gin",
	}
	engine.Use(
		server.CorsMiddleware, 
		server.ServerModelMiddleware,
		server.CheckRequestMiddleware,
	)
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

func (s Server) CorsMiddleware(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", s.allowedOrigin)
}

func (s Server) ServerModelMiddleware(c *gin.Context) {
	c.Header("X-Server-Model", s.model)
}

func (s Server) CheckRequestMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth != "ABC" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
}

func (s *Server) Run() error {
	return s.Engine.Run(fmt.Sprintf(":%d", s.port))
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

	
	c.JSON(200, categories)
}

func (s *Server) Products(c *gin.Context) {
	twoDollars := money.New(200, "USD")
	fourDollars := money.New(400, "USD")
	products := []product.Product{
		{
			ID:               "123",
			Name:             "Product1",
			Description:      "Product1 desc",
			ShortDescription: "ShortDesc1",
			PriceVATExcluded: product.Amount{Money: twoDollars, Display: twoDollars.Display()},
			VAT:              product.Amount{Money: twoDollars, Display: twoDollars.Display()},
			TotalPrice:       product.Amount{Money: fourDollars, Display: fourDollars.Display()},
			Image:            "https://conteudo.imguol.com.br/c/noticias/1c/2022/05/24/imagem-criada-no-imagen-prototipo-do-google-que-cria-imagens-baseadas-em-texto-neste-caso-um-cachorro-corgi-andando-de-bicicleta-na-times-square-usando-oculos-de-sol-e-chapeu-de-praia-1653397634334_v2_900x506.jpg",
		},
		{
			ID:               "124",
			Name:             "Product2",
			Description:      "Product2 desc",
			ShortDescription: "ShortDesc2",
			PriceVATExcluded: product.Amount{Money: twoDollars, Display: twoDollars.Display()},
			VAT:              product.Amount{Money: twoDollars, Display: twoDollars.Display()},
			TotalPrice:       product.Amount{Money: fourDollars, Display: fourDollars.Display()},
			Image:            "https://conteudo.imguol.com.br/c/noticias/1c/2022/05/24/imagem-criada-no-imagen-prototipo-do-google-que-cria-imagens-baseadas-em-texto-neste-caso-um-cachorro-corgi-andando-de-bicicleta-na-times-square-usando-oculos-de-sol-e-chapeu-de-praia-1653397634334_v2_900x506.jpg",
		},
		{
			ID:               "125",
			Name:             "Product3",
			Description:      "Product3 desc",
			ShortDescription: "ShortDesc3",
			PriceVATExcluded: product.Amount{Money: twoDollars, Display: twoDollars.Display()},
			VAT:              product.Amount{Money: twoDollars, Display: twoDollars.Display()},
			TotalPrice:       product.Amount{Money: fourDollars, Display: fourDollars.Display()},
			Image:            "https://conteudo.imguol.com.br/c/noticias/1c/2022/05/24/imagem-criada-no-imagen-prototipo-do-google-que-cria-imagens-baseadas-em-texto-neste-caso-um-cachorro-corgi-andando-de-bicicleta-na-times-square-usando-oculos-de-sol-e-chapeu-de-praia-1653397634334_v2_900x506.jpg",
		},
	}
	
	c.JSON(200, products)
}
