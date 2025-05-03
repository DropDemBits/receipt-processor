// All routes for the web service.
package routes

import (
	"github.com/gin-gonic/gin"

	"receipt-processor/controllers/receipt"
)

func Router() *gin.Engine {
	r := gin.Default()

	{
		receipts := r.Group("/receipts")
		receipts.POST("/process", receipt.Process)
		receipts.GET("/:id/points", receipt.GetPoints)
	}

	return r
}
