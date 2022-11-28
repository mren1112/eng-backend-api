package universityhandle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uh *universityAllHandler) GetUniversityAllUnicon(c *gin.Context) {

	resUniversityAll, err := uh.universityAll_Res.GetUniversityAllUnicon()
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}
	c.JSON(http.StatusOK, resUniversityAll)
}
