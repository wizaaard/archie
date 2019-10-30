package routes

import (
	"archie/utils"
	"archie/utils/configer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Serve() {
	config := configer.LoadServeConfig()

	router := gin.Default()
	//router.Use(cors.New(cors.Config{
	//	AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
	//	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowAllOrigins:  false,
	//	AllowOriginFunc:  func(origin string) bool { return true },
	//	MaxAge:           86400,
	//}))

	//router.Use(func(c *gin.Context) {
	//	for k, v := range c.Request.Header {
	//		fmt.Println("")
	//		fmt.Println(k, v)
	//		fmt.Println("")
	//	}
	//})

	userRouter(router)
	organizationRouter(router)
	DocRouter(router)
	TodoRouter(router)

	utils.Logger(fmt.Sprintf("Listing on %s", config.Port))
	router.Run(config.GetAddress())
}
