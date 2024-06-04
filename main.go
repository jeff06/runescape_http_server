package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"runescape_http_server/ent"
	"runescape_http_server/ent/skill"
	"strconv"
)

type Server struct {
	db   *ent.Client
	http *gin.Engine
}

type Response struct {
	Data interface{} `json:"data"`
}

var svr Server

func main() {
	initDatabase()
	runHttpServer()
}

func initDatabase() {
	// init PostgreSQL
	client, err := ent.Open("sqlite3", "./identifier.sqlite?_fk=1")

	// init MySQL
	//client, err := ent.Open("mysql", "root:123456@tcp(localhost:3306)/test?parseTime=True")
	// init SQLite
	//client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	// init Gremlin (AWS Neptune)
	//client, err := ent.Open("gremlin", "http://localhost:8182")
	if err != nil {
		log.Fatal(err)
		return
	}
	svr.db = client.Debug()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func runHttpServer() {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	svr.http = r

	// api doc http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Read
	r.GET("/skills/:id", getOneSkill)

	// Read
	r.GET("/skills", getAllSkills)

	// Listen and serve on 0.0.0.0:8080
	_ = r.Run(":8080")
}

// @Summary get user
// @Produce application/json
// @Param username path string true "username"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/{username} [get]
func getOneSkill(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		// ... handle error
		panic(err)
	}
	dbSkill, _ := svr.db.Skill.Query().Where(skill.ID(id)).All(context.Background())
	if dbSkill == nil {
		ResponseJSON(c, http.StatusNotFound, nil)
		return
	}

	ResponseJSON(c, http.StatusOK, dbSkill)
}

func getAllSkills(c *gin.Context) {
	dbSkill, _ := svr.db.Skill.Query().All(context.Background())
	if dbSkill == nil {
		ResponseJSON(c, http.StatusNotFound, nil)
		return
	}

	ResponseJSON(c, http.StatusOK, dbSkill)
}

func ResponseJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Response{
		Data: data,
	})
}
