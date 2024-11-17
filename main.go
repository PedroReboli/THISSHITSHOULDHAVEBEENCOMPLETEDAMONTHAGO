package main

import "server/api"
func main(){
	route := api.RegisterRoutes() 
	route.Run(":8080")
}