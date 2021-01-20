package admins

import (
	"fmt"
	"gin-admin/app/http"
	"gin-admin/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type ConfigController struct {
	http.BaseController
}

func (c ConfigController) Init(ctx *gin.Context) {
	config := `
server:
  port: 886
  domain: "http://127.0.0.1:8088"
  readTimeOut: "10"
  writeTimeOut: "10"
  maxHeaderBytes: "20"
  debug: "false"
  layout: "templates/layouts"
  upload: "static/upload"
mysql:
  debug: "true"
  host: "{host}"
  port: "{port}"
  username: "{name}"
  password: "{password}"
  db: "admin"
  charset: "utf8mb4"
jwt:
  signingKey: "tnjKJX2u6hhNA9M94ZVqJlwvWJImy6kDQAPhQamtpR9KxOjzORz75FhG7Vic8tsn"
  adminIdentityKey: "Admin"
  identityKey: "User"
  expiresAt: "30"
`
	form := struct {
		Port     int    `json:"port" form:"port" binging:"required"`
		Host     string `json:"host" form:"host" binging:"required"`
		Name     string `json:"name" form:"name" binging:"required"`
		DB       string `json:"db" form:"db" binging:"required"`
		Password string `json:"password" form:"password"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		lang := make(map[string]string)
		lang["Name"] = "用户名"
		lang["Password"] = "密码"
		lang["Port"] = "端口号"
		lang["Host"] = "数据库连接地址"
		err := c.Translate(err, lang)
		if err != nil {
			c.Error(ctx, err.Error())
		} else {
			c.Error(ctx, "")
		}
		return
	}
	config = strings.ReplaceAll(config, "{port}", fmt.Sprintf("%v", form.Port))
	config = strings.ReplaceAll(config, "{host}", form.Host)
	config = strings.ReplaceAll(config, "{db}", form.DB)
	config = strings.ReplaceAll(config, "{name}", form.Name)
	config = strings.ReplaceAll(config, "{password}", form.Password)
	var d1 = []byte(config)
	dir, _ := os.Getwd()
	mainFile := dir + "/main"
	envFile := dir + "/env.yml"
	_, err := os.Stat(mainFile)
	if err != nil {
		mainFile = "/Applications/react-desktop.app/Contents/Resources/main"
		envFile = "/Applications/react-desktop.app/Contents/Resources/env.yml"
	}
	err = ioutil.WriteFile(envFile, d1, 0666) //写入文件(字节数组)
	cmd := exec.Command(mainFile, "seed")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err == nil {
		c.Success(ctx, gin.H{})
	} else {
		c.Error(ctx, err.Error())
	}
}
func (c ConfigController) Check(ctx *gin.Context) {
	v := utils.Config()
	v.WatchConfig()
	db := v.GetString("mysql.db")
	v.OnConfigChange(func(e fsnotify.Event) {
		db = v.GetString("mysql.db")
	})
	if db == "" {
		c.Error(ctx, db)
	} else {
		c.Success(ctx, gin.H{})
	}
}
