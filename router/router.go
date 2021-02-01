package router

import (
	"gin-admin/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var engine *gin.Engine

func init() {
	binding.Validator = new(utils.DefaultValidator)
	gin.SetMode(gin.ReleaseMode)
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

var (
	g errgroup.Group
)

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Static("/", "./static")
	return e
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
	server02 := &http.Server{
		Addr:         ":889",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return s.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
