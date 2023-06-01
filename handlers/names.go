// handlers/names.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/AD-Singh-S-Barlow/KiraKillsRecord/models"
)

func GetNames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Names)
}

func PutNames(c *gin.Context) {
	var update models.NameList

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Names = append(models.Names, update)

	c.IndentedJSON(http.StatusCreated, update)
}

func PostNames(c *gin.Context) {
	var newName models.NameList

	if err := c.BindJSON(&newName); err != nil {
		return
	}

	models.Names = append(models.Names, newName)

	c.IndentedJSON(http.StatusCreated, newName)
}

func DeleteNames(c *gin.Context) {
	var deleteItem models.NameList

	if err := c.ShouldBindJSON(&deleteItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var remove bool
	for i, d := range models.Names {
		if d.NameF == deleteItem.NameF && d.NameL == deleteItem.NameL && d.Age == deleteItem.Age {
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
