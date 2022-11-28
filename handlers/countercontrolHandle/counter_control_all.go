package countercontrolhandle

import (
	//"eng-backend-api/services/countercontrolservices"
	"fmt"
	//"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uh *counterControlAllHandler) GetCounterControlAll(c *gin.Context) {

	counterControlAll, err := uh.countercontrolAll_Res.GetCounterControlAllData()
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}
	c.JSON(http.StatusOK, counterControlAll)
}
