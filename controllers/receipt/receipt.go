// REST endpoint for receipts.
package receipt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// post /process
//
// Submit a receipt to process.
// Returns a UUID to use with [GetPoints].
func Process(cx *gin.Context) {
	cx.Status(http.StatusBadRequest)
}

// get /:id/points
//
// Get the points awarded from a processed receipt.
func GetPoints(cx *gin.Context) {
	cx.Status(http.StatusNotFound)
}
