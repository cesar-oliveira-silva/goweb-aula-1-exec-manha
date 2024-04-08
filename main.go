package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	Id          int     `json:"id"`
	Nome        string  `json:"nome"`
	Sobrenome   string  `json:"sobrenome"`
	Email       string  `json:"email"`
	Idade       int     `json:"idade"`
	Altura      float64 `json:"altura"`
	Ativo       bool    `json:"ativo"`
	DataCriação string  `json:"dataCriacao"`
}

func main() {

	router := gin.Default()

	router.GET("/exercicio2", func(c *gin.Context) {
		c.JSON(200, gin.H{"mensagem": "Olá Cesar"})
	})

	router.GET("/listaUsuarios/GetAll", func(c *gin.Context) {

		var todos = []usuario{}

		todos = append(todos, usuario{
			Id:          1,
			Nome:        "cesar",
			Sobrenome:   "silva",
			Email:       "email",
			Idade:       15,
			Altura:      1.82,
			Ativo:       true,
			DataCriação: "2024-04-08"})

		todos = append(todos, usuario{
			Id:          2,
			Nome:        "augusto",
			Sobrenome:   "silva",
			Email:       "email",
			Idade:       15,
			Altura:      1.82,
			Ativo:       true,
			DataCriação: "2024-04-08"})

		fmt.Println("lista todos: ", todos)

		// jsonData, err := json.Marshal(todos)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Printf("json gerado: %s \n", jsonData)
		c.JSON(200, todos)

	})

	router.Run()

}
