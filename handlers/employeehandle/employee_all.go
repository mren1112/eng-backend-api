package employeehandle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uh *employeeAllHandler) GetEmployeeAllUnicon(c *gin.Context) {

	resEmployeeAll, err := uh.employeeAll_Res.GetEmployeeAllUnicon()
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}
	c.JSON(http.StatusOK, resEmployeeAll)
}
