package facultyHandle

import (
	"eng-backend-api/services/facultyServices"
	"fmt" 
 
	//"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uh *facultyAllHandler) GetFacultyAll(c *gin.Context) {

	resfacultyAll, err := uh.facultyAll_Res.GetFacultyAllData()
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}
	c.JSON(http.StatusOK, resfacultyAll)
}

func (uh *facultyAllHandler) GetFacultyById(c *gin.Context) {
	param1 := c.Param("param1")
	//fmt.Println(param1) 
	resfacultyAll, err := uh.facultyAll_Res.GetFacultyById(param1)
	if err != nil {
		
    c.String(405, "method not allowed")
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}
	c.JSON(http.StatusOK, resfacultyAll)
}

 

func (uh *facultyAllHandler) GetFacultyByJS(c *gin.Context) {
	var requestBoby facultyServices.FacultyAllRespone 
	//var requestBoby Faculty_body 
	var jsonArr = []string{}
	
 
	err := c.ShouldBindJSON(&requestBoby)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}

	jsonArr = append(jsonArr, requestBoby.FACULTY_NO)
	jsonArr = append(jsonArr, requestBoby.MAJOR_NO)

	fmt.Println("res ", requestBoby)  

	resfacultyAll, err := uh.facultyAll_Res.GetFacultyByJS(jsonArr)
	if err != nil {
		
    c.String(405, "method not allowed")
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}

	 
	c.JSON(http.StatusOK, resfacultyAll)
}


/* func (uh *facultyAllHandler) GetFacultyById(c *gin.Context) {

	paramID := c.Param("id")
	resfacultyAll, err := uh.facultyAll_Res.GetFacultyById(paramID)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "status": "400"})
		return
	}
	c.JSON(http.StatusOK, resfacultyAll)
}

*/
