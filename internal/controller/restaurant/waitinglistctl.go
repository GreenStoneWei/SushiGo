package restaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	waitinglistModel "stone.com/sushigo/internal/model/waitinglist"
)

func PatchWaitingListById(c *gin.Context) {
	waitingListId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	var requestBody waitinglistModel.PatchWaitingListRequest
	err = c.BindJSON(&requestBody)
	if err != nil {
		panic(err)
	}

	waitinglistModel.PatchWaitingListById(waitingListId, requestBody)

	c.IndentedJSON(http.StatusOK, gin.H{"result": "sucess"})
}
