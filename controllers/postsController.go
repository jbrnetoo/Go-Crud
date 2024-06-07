package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

// Endpoint para inserir Posts
// @Summary Inseri um novo post na base
// @Description Inseri um novo post na base de dados
// @Accept json
// @Produce json
// @Success 200 {integer} integer
// @Router /posts [post]
func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	// faz o bind do body da requisição para o struct
	err := c.Bind(&body)

	// verifica por erro no Bind
	if err != nil {
		println(err.Error())
	}

	// monta o objeto
	post := models.Post{Title: body.Title, Body: body.Body}

	// insere na base
	result := initializers.DB.Create(&post)

	// verifica por erros
	if result.Error != nil {
		c.Status(400)
		return
	}

	// retorno do endpoint
	c.JSON(200, gin.H{
		"post": post,
	})
}

// Endpoint para obter todos os posts
// @Summary Consulta todos os post na base
// @Description Consulta todos os post na base de dados
// @Accept json
// @Produce json
// @Success 200 {integer} integer
// @Router /posts [get]
func PostsIndex(c *gin.Context) {
	// pega todos os posts da base
	var posts []models.Post
	initializers.DB.Find(&posts)

	// retorno do endpoint
	c.JSON(200, gin.H{
		"post": posts,
	})
}

// Endpoint para obter post por ID
// @Summary Consulta post pelo ID na base
// @Description Consulta post pelo ID na base
// @Accept json
// @Produce json
// @Param id query int true "Número 1"
// @Success 200 {integer} integer
// @Router /posts:id [get]
func PostsShow(c *gin.Context) {
	// obtem parâmetro da rota
	id := c.Param("id")

	// obtem post da base
	var post models.Post
	initializers.DB.First(&post, id)

	// retorno do endpoint
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	// obtem parâmetro da rota
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	// faz o bind do body da requisição para o struct
	err := c.Bind(&body)

	// verifica por erro no Bind
	if err != nil {
		println(err.Error())
	}

	// obtem o Post da base
	var post models.Post
	initializers.DB.First(&post, id)

	// faz o update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// retorno do endpoint
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	// obtem parâmetro da rota
	id := c.Param("id")

	// deleta da base
	initializers.DB.Delete(&models.Post{}, id)

	// retorno do endpoint
	c.Status(204)
}
