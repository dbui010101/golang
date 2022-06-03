package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"fmt"
	"net/http"
)
var router *gin.Engine


var listStroke [] string
//var birdData map[string]interface{}
//var *data map[string]interface{}
func main() {
	go h.run()
	router := gin.New()
	router.LoadHTMLGlob("templates/*")
  router.StaticFile("/bg.jpg", "./assets/bg.jpg")
  router.StaticFile("/crayon.svg", "./assets/crayon.svg")
  router.StaticFile("/admin.svg", "./assets/admin.svg")

	router.GET("/room/:roomId", func(c *gin.Context) {
		fmt.Print("test")
    fmt.Println("nouveau membre")
		c.HTML(200, "salon.html", gin.H{
			"pseudo_courant": "jamie",
			"pseudo": listeUtilisateur,
			"data": &listStroke,
		})
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"salons": listeSalon,
		})
	})

  router.Use(cors.Default())

	router.GET("/salons/:nomRoom/:id/:pseudo", func(c *gin.Context) {
		nomRoom := c.Param("nomRoom")
		id := c.Param("id")
		pseudo := c.Param("pseudo")
		envoieSalonDispo(c.Writer, c.Request, nomRoom, id)
		listeUtilisateur = append(listeUtilisateur, pseudo) //Slice will become [1, 2, 3, 4]

		// fmt.Println(listeUtilisateur)
	})

	router.GET("/add/:pseudo/:room", func(c *gin.Context) {
		room := c.Param("room")
		pseudo := c.Param("pseudo")
		ajoutPseudo(c.Writer, c.Request, room, pseudo)
	})

	router.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"joueurs": listeUtilisateur,
		})
	})

	router.Run("0.0.0.0:8080")

	
}
