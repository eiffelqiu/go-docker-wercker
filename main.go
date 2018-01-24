package main

import (  
    "log"
    "github.com/gin-gonic/gin"
)

func main() { 
	app := gin.Default()
	app.GET("/", IndexPage)	
	app.Run(":80")
} 

func checkErr(err error, exp string) {
	if (err != nil) {
		log.Fatalf("Error -> %v", err.Error())
	}
}

func setHeaders(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

    if c.Request.Method == "OPTIONS" {
        return
    }	
}

func IndexPage(c *gin.Context) {
	setHeaders(c)
	content := gin.H{}
	content["data"] = []string{"Chocolate","Coffee","Bread"}
	c.JSON(200, content)
}