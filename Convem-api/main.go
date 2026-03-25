package main //Onde a API será executada

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	ID      string `json:"id"`
	Holder  string `json:"holder"`
	Balance int64  `json:"balance"`
}

var accounts = []Account{
	{ID: "1", Holder: "Candidato Convem", Balance: 100000},
	{ID: "2", Holder: "Empresa Parceira", Balance: 50000},
}

func getBalance(context *gin.Context) {
	id := context.Param("id") //Pega o ID da URL e mostra
	for _, acc := range accounts {
		if acc.ID == id {
			//Se achar retorna a conta correspondente a busca...
			context.IndentedJSON(http.StatusOK, acc)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Conta não encontrada"})
}

func transferMoney(context *gin.Context) {
	var request struct {
		FromID string `json:"from_id"`
		ToID   string `json:"to_id"`
		Amount int64  `json:"amount"`
	}

	if err := context.BindJSON(&request); err != nil {
		return
	}

	for i := range accounts {
		if accounts[i].ID == request.FromID {
			accounts[i].Balance -= request.Amount //Manda o dinheiro
		}
		if accounts[i].ID == request.ToID {
			accounts[i].Balance += request.Amount //Recebe o dinheiro
		}
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Transferência realizada com sucesso!"})
}

func main() {
	router := gin.Default()

	router.GET("/accounts/:id/balance", getBalance)
	router.POST("/transfer", transferMoney)
	router.Run("localhost:8080")
}
