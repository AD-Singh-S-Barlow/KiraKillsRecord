
package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/AD-Singh-S-Barlow/KiraKillsRecord/models"
)
//retrieves all the names
func GetNames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Names)
}
//adds a new name to the list
func PutNames(c *gin.Context) {
	var update models.NameList

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
//Append the 'update' to the 'models.Names' slice
	models.Names = append(models.Names, update)

	c.IndentedJSON(http.StatusCreated, update)
}
//updates an existing name in the list
func PostNames(c *gin.Context) {
	var newName models.NameList

// Bind the JSON payload to the 'newName' variable
	if err := c.BindJSON(&newName); err != nil {
		return
	}
// Append the 'newName' to the 'models.Names' slice
	models.Names = append(models.Names, newName)

	c.IndentedJSON(http.StatusCreated, newName)
}
//removes a name from the list
func DeleteNames(c *gin.Context) {
	var deleteItem models.NameList
// Bind the JSON payload to the 'deleteItem' variable
	if err := c.ShouldBindJSON(&deleteItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var remove bool
	for i, d := range models.Names {
		if d.NameF == deleteItem.NameF && d.NameL == deleteItem.NameL && d.Age == deleteItem.Age {
			// Remove the item from the 'models.Names' slice
			models.Names = append(models.Names[:i], models.Names[i+1:]...)
			remove = true
			break
		}
	}

	if remove {
		c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
	}
}
