package api

import (
	"math/rand"
	"net/http"
	"server/dao"
	"server/dao/model"

	"github.com/gin-gonic/gin"
)
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8000")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func IsLogged(callbackFunc func(*gin.Context, *model.User)) gin.HandlerFunc{
	return func(c *gin.Context) {
		token, ok := c.Request.Header[http.CanonicalHeaderKey("token")]
		if !ok {
			c.Status(401)
			return
		}
		var player model.User
		dao.DB.Where(&model.User{Token: token[0]}).First(&player)
		callbackFunc(c, &player)
	}
}
func RequireRole(callbackFunc func(*gin.Context, *model.User), role int32) func(ctx *gin.Context, usr *model.User){
	return func(ctx *gin.Context, usr *model.User){
		if usr.Role != role{
			ctx.Status(401)
			return
		}
		callbackFunc(ctx, usr)
	}
}
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

type LoginDto struct{
	User string `json:"User"`
	Password string `json:"Password"`
}

func Login(ctx *gin.Context){
	// var loginObj LoginDto
	// err := ctx.BindJSON(&loginObj)
	// if err != nil{
	// 	println(err.Error())
	// 	println("bd login object")
	// 	ctx.AbortWithStatus(400)
	// 	return
	// }
	// cliente := model.User{Nome: loginObj.User, Senha: loginObj.Password}
	// // dao.DB.WithContext(dao.DB.Statement.Context).PrepareStmt
	// out := dao.DB.Model(&cliente).First(&cliente)
	// if out.Error != nil{
	// 	ctx.AbortWithStatus(400)
	// }
	// cliente.TokenAtual = randSeq(25)
	// // dao.DB.Model(&cliente).Where("id = ?", cliente.ID).Update("token_atual")
	// out = dao.DB.Save(&cliente)
	
	// if out.Error != nil{
	// 	ctx.AbortWithStatus(500)
	// }

	// // ctx.SetCookie("credentials",cliente.TokenAtual, -1, "/", "*", false,false)

	// ctx.Data(200,"application/json; charset=utf-8", []byte("\""+cliente.TokenAtual+"\""))
	
}

func RegisterRoutes() *gin.Engine{
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/update", IsLogged(RequireRole(CreateEvent, 1)))
	return r
}