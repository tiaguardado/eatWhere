package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

// Estrutura do prato
type Prato struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Nome        string `json:"nome"`
	Tag         string `json:"tag"`
	Restaurante string `json:"restaurante"`
}

// Variável global do banco de dados
var db *gorm.DB

func initDB() {
	var err error
	dsn := "host=localhost user=postgres password=1234 dbname=pratos_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco de dados")
	}

	// Criar a tabela se não existir
	db.AutoMigrate(&Prato{})

	// Inserir dados iniciais
	var pratos = []Prato{
		{Nome: "Bife à Portuguesa", Tag: "Carne", Restaurante: "Cantinho do Sabor"},
		{Nome: "Sopa de Legumes", Tag: "Sopa", Restaurante: "Aconchego da Vó"},
		{Nome: "Frango Grelhado", Tag: "Carne", Restaurante: "Churrascaria do Bairro"},
		{Nome: "Pizza Margherita", Tag: "Pizza", Restaurante: "Pizzaria Bella Roma"},
		{Nome: "Bacalhau à Brás", Tag: "Bacalhau", Restaurante: "Taberna do Bacalhau"},
		{Nome: "Risoto de Cogumelos", Tag: "Arroz", Restaurante: "Restaurante La Dolce Vita"},
		{Nome: "Churrasco", Tag: "Carne", Restaurante: "Churrascaria Rio Grande"},
		{Nome: "Sardinhas Grelhadas", Tag: "Peixe", Restaurante: "Marisqueira do Porto"},
		{Nome: "Polvo à Lagareiro", Tag: "Marisco", Restaurante: "O Cantinho do Mar"},
		{Nome: "Feijoada", Tag: "Tradicional", Restaurante: "Casa da Feijoada"},
		{Nome: "Salada de Frutos do Mar", Tag: "Marisco", Restaurante: "Oceano Fresco"},
		{Nome: "Francesinha", Tag: "Sanduíche", Restaurante: "Restaurante Sabor e Tradição"},
		{Nome: "Arroz de Pato", Tag: "Arroz", Restaurante: "Taberna do Pato Real"},
		{Nome: "Cozido à Portuguesa", Tag: "Tradicional", Restaurante: "Restaurante Típico Português"},
		{Nome: "Sopa de Cebola", Tag: "Sopa", Restaurante: "Aldeia dos Sabores"},
		{Nome: "Massa à Carbonara", Tag: "Massa", Restaurante: "Trattoria da Nonna"},
		{Nome: "Bacalhau com Natas", Tag: "Bacalhau", Restaurante: "Bacalhau & Companhia"},
		{Nome: "Frango Assado", Tag: "Carne", Restaurante: "Churrasco de Ouro"},
		{Nome: "Espaguete à Bolonhesa", Tag: "Massa", Restaurante: "La Spaghettata"},
		{Nome: "Picanha na Brasa", Tag: "Carne", Restaurante: "Picanha Grill"},
		{Nome: "Arroz de Marisco", Tag: "Marisco", Restaurante: "Mariscada & Cia"},
		{Nome: "Feijão Tropeiro", Tag: "Tradicional", Restaurante: "Cozinha do Sertão"},
		{Nome: "Escondidinho de Carne Seca", Tag: "Tradicional", Restaurante: "Sabor Nordestino"},
		{Nome: "Bacalhau com Broa", Tag: "Bacalhau", Restaurante: "Maré Alta"},
		{Nome: "Tarte de Maçã", Tag: "Sobremesa", Restaurante: "Doce Encanto"},
		{Nome: "Sopa de Peixe", Tag: "Sopa", Restaurante: "O Mar da Sopa"},
		{Nome: "Hambúrguer Gourmet", Tag: "Sanduíche", Restaurante: "Burger House"},
		{Nome: "Pizza Quatro Queijos", Tag: "Pizza", Restaurante: "Pizzaria da Praça"},
		{Nome: "Lasanha à Bolonhesa", Tag: "Massa", Restaurante: "Cantina do Nonno"},
		{Nome: "Mocotó", Tag: "Tradicional", Restaurante: "Restaurante do Sertão"},
		{Nome: "Peixe Assado", Tag: "Peixe", Restaurante: "Marisqueira do Atlântico"},
		{Nome: "Arroz de Marisco", Tag: "Marisco", Restaurante: "Marisqueira do Porto"},
		{Nome: "Frango Frito", Tag: "Carne", Restaurante: "Churrascaria do Bairro"},
		{Nome: "Pizza Calabresa", Tag: "Pizza", Restaurante: "Pizzaria Bella Roma"},
		{Nome: "Risoto de Frutos do Mar", Tag: "Arroz", Restaurante: "Restaurante La Dolce Vita"},
		{Nome: "Bacalhau à Lagareiro", Tag: "Bacalhau", Restaurante: "Taberna do Bacalhau"},
		{Nome: "Feijoada", Tag: "Tradicional", Restaurante: "Casa da Feijoada"},
		{Nome: "Churrasco Misto", Tag: "Carne", Restaurante: "Churrascaria Rio Grande"},
		{Nome: "Tábua de Frios", Tag: "Petisco", Restaurante: "Restaurante do Sertão"},
		{Nome: "Camarão ao Alho e Óleo", Tag: "Marisco", Restaurante: "Marisqueira do Atlântico"},
		{Nome: "Massa de Peixe", Tag: "Massa", Restaurante: "Pizzaria Bella Roma"},
		{Nome: "Picanha no Sal Grosso", Tag: "Carne", Restaurante: "Churrascaria Rio Grande"},
		{Nome: "Pão de Queijo", Tag: "Petisco", Restaurante: "Aconchego da Vó"},
		{Nome: "Lulas Grelhadas", Tag: "Peixe", Restaurante: "Marisqueira do Porto"},
		{Nome: "Camarão na Moranga", Tag: "Marisco", Restaurante: "O Cantinho do Mar"},
		{Nome: "Salada Caesar", Tag: "Salada", Restaurante: "La Spaghettata"},
		{Nome: "Bacalhau à Zé do Pipo", Tag: "Bacalhau", Restaurante: "Bacalhau & Companhia"},
		{Nome: "Alheira", Tag: "Tradicional", Restaurante: "Restaurante Típico Português"},
		{Nome: "Carne de Panela", Tag: "Carne", Restaurante: "Restaurante Sabor e Tradição"},
	}

	for _, prato := range pratos {
		db.FirstOrCreate(&prato, Prato{Nome: prato.Nome})
	}
}

func main() {
	initDB()

	r := gin.Default()

	// Rota para listar todos os pratos
	r.GET("/pratos", func(c *gin.Context) {
		var pratos []Prato
		db.Find(&pratos)
		c.JSON(http.StatusOK, pratos)
	})

	// Rota para listar pratos por tag
	r.GET("/pratos/:tag", func(c *gin.Context) {
		tag := c.Param("tag")
		var filtrados []Prato
		db.Where("tag = ?", tag).Find(&filtrados)
		c.JSON(http.StatusOK, filtrados)
	})

	r.GET("/restaurantes/:tag1/:tag2", func(c *gin.Context) {
		tag1 := c.Param("tag1")
		tag2 := c.Param("tag2")

		var pratos []Prato
		// Busca os pratos que têm a tag "tag1" ou "tag2"
		db.Where("tag = ? OR tag = ?", tag1, tag2).Find(&pratos)

		// Criar um mapa para armazenar os pratos de cada restaurante
		restauranteMap := make(map[string]map[string][]Prato)

		for _, prato := range pratos {
			if restauranteMap[prato.Restaurante] == nil {
				restauranteMap[prato.Restaurante] = make(map[string][]Prato)
			}
			// Adiciona o prato à lista do restaurante correspondente, agrupando por tag
			restauranteMap[prato.Restaurante][prato.Tag] = append(restauranteMap[prato.Restaurante][prato.Tag], prato)
		}

		// Agora, verificar se algum restaurante tem ambos os pratos (tag1 e tag2)
		var restaurantesComAmbasTags []map[string]interface{}
		for restaurante, tags := range restauranteMap {
			// Verifica se o restaurante tem pratos de ambas as tags
			hasTag1 := tags[tag1] != nil
			hasTag2 := tags[tag2] != nil

			if hasTag1 && hasTag2 {
				// Adiciona o restaurante e seus pratos às duas tags
				restaurantesComAmbasTags = append(restaurantesComAmbasTags, map[string]interface{}{
					"restaurante": restaurante,
					"pratos":      tags,
				})
			}
		}

		// Se nenhum restaurante foi encontrado
		if len(restaurantesComAmbasTags) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Nenhum restaurante encontrado com pratos de ambas as tags"})
		} else {
			c.JSON(http.StatusOK, restaurantesComAmbasTags)
		}
	})

	// Rota para listar pratos por restaurante
	r.GET("/pratos/restaurante/:restaurante", func(c *gin.Context) {
		restaurante := c.Param("restaurante")
		var filtrados []Prato
		db.Where("restaurante = ?", restaurante).Find(&filtrados)
		c.JSON(http.StatusOK, filtrados)
	})

	// Rota para adicionar um prato à lista
	r.POST("/pratos", func(c *gin.Context) {
		var prato Prato
		if err := c.ShouldBindJSON(&prato); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&prato)
		c.JSON(http.StatusCreated, prato)
	})

	// Rota para consumir uma API externa usando Resty
	r.GET("/api-externa", func(c *gin.Context) {
		client := resty.New()
		resp, err := client.R().Get("https://api.exemplo.com/dados")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": resp.String()})
	})

	// Iniciar servidor
	r.Run(":8080")
}

//http://localhost:8080/pratos/
/* {
    "nome": "Feijoada de Leitão",
    "tag": "Leitão",
    "restaurante": "Pote"
}*/

// http://localhost:8080/pratos/restaurante/Restaurante A
