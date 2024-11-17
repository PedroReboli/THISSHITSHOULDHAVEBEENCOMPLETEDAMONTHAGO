package api

import (
	"time"

	"github.com/gin-gonic/gin"
)


func InviteVolunteer(ctx *gin.Context){

}

func CancelInvite(ctx *gin.Context){

}

func KickVolunteer(ctx *gin.Context){

}

func AcceptApplication(ctx *gin.Context){

}

func RejectApplication(ctx *gin.Context){

}

type Manager struct{
	Id int
	Name string
	BirthDay time.Time
	Qualifications string
	Photo string
	Bio string
}

type Schedule struct{
	Id int
	EventId int
	ManagerId int
	VolunteerId int
	StartDate time.Time
	FinishDate time.Time
	Weakly bool

}