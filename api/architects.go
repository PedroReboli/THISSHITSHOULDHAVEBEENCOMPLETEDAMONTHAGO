package api

import (
	"server/dao"
	"server/dao/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEvent(ctx *gin.Context){

}

func CreateEvent(ctx *gin.Context, user *model.User){
	var event Event
	err := ctx.BindJSON(&event)
	if err != nil{
		ctx.AbortWithError(400, err)
		return
	}
	dbEvent := model.Event{Name: event.Name,
		Details: event.Details,
		Architect: user.ID,
		Begindate: event.Date}
	tx := dao.DB.Create(dbEvent)
	if tx.Error != nil{
		ctx.AbortWithError(500, err)
		return
	}
}

func EditEvent(ctx *gin.Context, user *model.User){
	
}

func RemoveEvent(ctx *gin.Context, user *model.User){
	StringId := ctx.Params.ByName("id")
	id, err := strconv.Atoi(StringId)
	if err != nil{
		ctx.AbortWithError(500, err)
		return
	}
	dao.DB.Delete(&model.Event{ID: int32(id), Architect: user.ID})
}

func AddManager(ctx *gin.Context){

}

func RemoveManager(ctx *gin.Context){

}

type Event struct{
	Name string `json:"Name"`
	Date time.Time `json:"Date"`
	Details string `json:"Details"`
}

// type Event struct{
// 	Name string `json:"Name"`
// 	Date time.Time `json:"Date"`
// 	Details string `json:"Details"`
// }

type EventManager struct{
	ManagerId int
	EventId int
}