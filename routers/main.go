package routers

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func apiVersion(orig string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	prefix:= os.Getenv("APP_PREFIX")
	return fmt.Sprintf("%s%s", prefix, orig)
}

func RouterCombain(r *gin.Engine) {

	RolesRouters(r.Group(apiVersion("/roles")))
	UserRouters(r.Group(apiVersion("/users")))
	CategoriesRouters(r.Group(apiVersion("categories")))
	ProductsRouters(r.Group(apiVersion("/products")))
	AuthRouters(r.Group(apiVersion("/auth")))
	ProfileRouters(r.Group(apiVersion("/profile")))
	PromoRouters(r.Group(apiVersion("/promo")))
	TestimonialsRouters(r.Group(apiVersion("/testimonials")))
	OrderTypesRouters(r.Group(apiVersion("/order-type")))
	CartsRouters(r.Group(apiVersion("/carts")))
	TransactionRouters(r.Group(apiVersion("/transaction")))
	TransactionStatusRouters(r.Group(apiVersion("/transaction-status")))
}
