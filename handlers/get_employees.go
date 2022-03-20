package employees

import (
	"context"
	"time"
emp "github.com/tusharhow/go-oms/models"
db "github.com/tusharhow/go-oms/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEmployee(c *gin.Context) {
	query := bson.D{{}}
	collection := db.MGI.Db.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, query)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var results []*emp.Employee
	for cur.Next(ctx) {
		var elem emp.Employee
		err := cur.Decode(&elem)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		results = append(results, &elem)
	}
	c.JSON(200, results)

}