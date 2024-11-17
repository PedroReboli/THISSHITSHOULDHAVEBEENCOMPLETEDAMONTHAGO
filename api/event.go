package api

import (
	"server/dao"

	"github.com/gin-gonic/gin"
)


func SearchEvent(ctx *gin.Context){

}

func ListEvent(ctx *gin.Context){
	var events []Event
	tx := dao.DB.Find(&events)
	if tx.Error != nil{
		ctx.AbortWithError(500, tx.Error)
		return
	}
	ctx.AbortWithStatusJSON(200, events)
}