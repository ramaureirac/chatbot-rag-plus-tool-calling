package router

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	agent "github.com/ramaureirac/devops-ragbot/src/pkg/agent"
	public "github.com/ramaureirac/devops-ragbot/src/public"
)

type RouterRequest struct {
	Message string `json:"message" binding:"required"`
}

func NewRouterApp() *gin.Engine {
	switch os.Getenv("GIN_MODE") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "testing":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	sub, err := fs.Sub(public.GetWWWEmbed(), "dist/assets")
	if err != nil {
		log.Fatal("server: " + err.Error())
	}

	router := gin.New()
	router.SetTrustedProxies([]string{"192.168.1.0/24"})
	sess := Sessions{Agents: map[string]*agent.Agent{}}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Anon-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.StaticFS("/assets", http.FS(sub))
	//router.Static("/assets", "./src/server/dist")
	//router.LoadHTMLGlob("./src/dist/index.html")
	tmpl := template.Must(template.New("").ParseFS(public.GetWWWEmbed(), "dist/index.html"))
	router.SetHTMLTemplate(tmpl)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/login", func(c *gin.Context) {
		id, err := sess.RegisterSession()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": id,
		})
	})

	router.POST("/ask", xAnonIDMiddleware(&sess), func(c *gin.Context) {
		var req RouterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request to api"})
			return
		}
		res, err := sess.AskQuestion(c.GetHeader("X-Anon-ID"), req.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": res,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": res,
		})
	})

	tmout, err := strconv.ParseFloat(os.Getenv("SESSION_TIMEOUT"), 64)
	if err != nil {
		tmout = 5.0
	}

	sess.DropSessionsOlderThan(tmout)
	return router
}
