package employees



import (
	"context"
	"time"
	emp "github.com/tusharhow/go-oms/models"
	db "github.com/tusharhow/go-oms/db"
	"github.com/gin-gonic/gin"
)


func AddEmployee(c *gin.Context) {
	employee := new(emp.Employee)

	if err := c.ShouldBindJSON(employee); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	collection := db.MGI.Db.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, employee)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Employee added")

}