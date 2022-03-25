package employees

import (
	"context"
	"time"
	db "github.com/tusharhow/go-oms/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)



func DeleteEmployee(c *gin.Context) {
	idParam := c.Param("id")
	collection := db.MGI.Db.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.D{{
		
	}})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": idParam, "status": "Deleted"})

}


