// REST endpoint for receipts.
package receipt

import (
	"net/http"
	"receipt-processor/model"
	"receipt-processor/processor"

	"github.com/gin-gonic/gin"
)

var ruleProcessor = processor.DefaultRules()

// post /process
//
// Submit a receipt to process.
// Returns a UUID to use with [GetPoints].
func Process(cx *gin.Context) {
	var receipt model.Receipt

	if err := cx.ShouldBindBodyWithJSON(&receipt); err != nil {
		cx.Status(http.StatusBadRequest)
		return
	}

	if err := receipt.Validate(); err != nil {
		cx.Status(http.StatusBadRequest)
		return
	}

	points, err := ruleProcessor.Process(&receipt)
	if err != nil {
		cx.Status(http.StatusBadRequest)
		return
	}

	id := model.AddProcessedReceipt(model.ProcessedReceipt{Points: points})

	cx.JSON(http.StatusOK, gin.H{"id": id})
}

// get /:id/points
//
// Get the points awarded from a processed receipt.
func GetPoints(cx *gin.Context) {
	id := cx.Param("id")

	if processed, ok := model.GetProcessedReceipt(id); ok {
		cx.JSON(http.StatusOK, gin.H{"points": processed.Points})
	} else {
		cx.Status(http.StatusNotFound)
	}
}
