package studentshandle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uh *studentAllHandler) GetStudentAllUnicon(c *gin.Context) {

	resStudentAll, err := uh.studentAll_Res.GetStudentAllUnicon()
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}
	c.JSON(http.StatusOK, resStudentAll)
}

func (uh *studentAllHandler) GetStudentLimitUnicon(c *gin.Context) {
	param1 := c.Param("param1")

	resStudentAll, err := uh.studentAll_Res.GetStudentLimitUnicon(param1)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}
	c.JSON(http.StatusOK, resStudentAll)
}
