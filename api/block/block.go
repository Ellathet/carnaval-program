package block

import (
	"net/http"

	"github.com/Ellathet/carnaval/models"
	"github.com/Ellathet/carnaval/utils"
	"github.com/gin-gonic/gin"
)

func CreateBlock(c *gin.Context) {
	var block models.Block

	if err := c.ShouldBindJSON(&block); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.DB.Create(&block).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, block)
}

func GetBlock(c *gin.Context) {
	id := c.Param("id")
	var block models.Block

	if err := utils.DB.Where("id = ?", id).First(&block).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
		return
	}

	c.JSON(http.StatusOK, block)
}

func UpdateBlock(c *gin.Context) {
	id := c.Param("id")
	var block models.Block

	// If repository exists, or services, can reuse this function, but...
	if err := utils.DB.Where("id = ?", id).First(&block).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
		return
	}

	var newBlock models.Block
	if err := c.ShouldBindJSON(&newBlock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.DB.Model(&block).Updates(newBlock)
	c.JSON(http.StatusOK, block)
}

func DeleteBlock(c *gin.Context) {
	id := c.Param("id")
	var block models.Block

	if err := utils.DB.Where("id = ?", id).First(&block).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
		return
	}

	utils.DB.Where("id = ?", id).Delete(&block)
	c.JSON(http.StatusOK, gin.H{"message": "Block deleted successfully"})
}
