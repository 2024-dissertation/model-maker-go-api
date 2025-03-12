package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	services "github.com/Soup666/diss-api/services"
	"github.com/gin-gonic/gin"
)

type VisionController struct {
	VisionService services.VisionService
}

func NewVisionController(visionService services.VisionService) *VisionController {
	return &VisionController{VisionService: visionService}
}

func (c *VisionController) AnalyzeImage(ctx *gin.Context) {

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}
	defer file.Close()

	f, err := os.CreateTemp("", "sample")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}

	fmt.Println("Temp file name:", f.Name())

	_, err = io.Copy(f, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}

	image := fmt.Sprintf("/%s", f.Name())

	defer os.Remove(f.Name())

	result, err := c.VisionService.AnalyseImage(image)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to analyze the image"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
