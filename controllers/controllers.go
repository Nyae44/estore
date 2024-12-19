package controllers

import (
	"github.com/gin-gonic/gin"
)

func HashPassword(password string) string {
	return password
}
func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func Login() gin.HandlerFunc {

}

func SignUp() gin.HandlerFunc {

}

func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {

}
