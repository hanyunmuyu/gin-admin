package router

import (
	_ "gin-admin/docs"
	"gin-admin/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	url := ginSwagger.URL("http://localhost:886/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
	_ = s.ListenAndServe()
}
