package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var replicaId = newReplicaId()

func newReplicaId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)
}

func main() {
	e := echo.New()
	e.Use(mw.Logger())
	e.Get("/", getProducts)
	e.Run(":8000")
}

func getProducts(c *echo.Context) error {
	return c.JSON(http.StatusOK, ProductsResp{
		Products: []string{"apple", "banana"},
		Replica:  replicaId,
	})
}

type ProductsResp struct {
	Products []string `json:"products"`
	Replica  int      `json:"_replica"`
}
