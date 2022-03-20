package main

import (
	"log"
	"github.com/gin-gonic/gin"
	emp "github.com/tusharhow/go-oms/handlers"
	db "github.com/tusharhow/go-oms/db"
)




func main() {
	r := gin.Default()
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	r.POST("/addEmployee", emp.AddEmployee)
	r.GET("/getEmployee", emp.GetEmployee)
	r.PUT("/updateEmployee", emp.UpdateEmployee)
	r.DELETE("/deleteEmployee", emp.DeleteEmployee)
	log.Fatal(r.Run(":8080"))



}



