package router

import (
	"fmt"
	"gin-admin/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

var engine *gin.Engine

func init() {
	binding.Validator = new(utils.DefaultValidator)
	engine = gin.Default()
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	defaultConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	defaultConfig.ExposeHeaders = []string{"Content-Length"}
	engine.Use(cors.New(defaultConfig))
}

// router router
func router() *gin.Engine {
	return engine
}
func Run() {
	//taskRouter()
	//apiRouter()
	//webRouter()
	adminRouter()
	s := &http.Server{
		Addr:           ":886",
		Handler:        router(),
		ReadTimeout:    time.Duration(5) * time.Second,
		WriteTimeout:   time.Duration(5) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
