package employees

import (
	"context"
	"time"
	emp "github.com/tusharhow/go-oms/models"
	db "github.com/tusharhow/go-oms/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateEmployee(c *gin.Context) {
	idParam := c.Param("id")
	employee := new(emp.Employee)
	if err := c.ShouldBindJSON(employee); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	collection := db.MGI.Db.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.UpdateOne(ctx, bson.D{{}}, bson.D{{"$set", employee}})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	employee.ID = idParam
	c.JSON(200, "Employee updated")

}