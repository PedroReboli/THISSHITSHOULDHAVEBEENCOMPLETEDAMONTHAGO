package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RejectInvite(ctx *gin.Context){

}

func AcceptInvite(ctx *gin.Context){

}

func ApplyEvent(ctx *gin.Context){

}

func CancelApplyEvent(ctx *gin.Context){

}

func ExitEvent(ctx *gin.Context){

}

func ListVolunteers(ctx *gin.Context){

}

func SearchVolunteers(ctx *gin.Context){

}

type Volunteer struct{
	Id int
	Name string
	BirthDay time.Time
	Qualifications string
	Photo string
	Bio string
}
