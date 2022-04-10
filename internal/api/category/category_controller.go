package category

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/category"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	CategoryService *category.Service
}

func NewCategoryController(categoryService *category.Service) *Controller {
	return &Controller{
		CategoryService: categoryService,
	}
}

// CreateCategory creates reads csv file and creates new categories in the database
func (c Controller) CreateCategory(context *gin.Context) {

	file, err := context.FormFile("file")

	if err != nil {
		context.JSON(400, gin.H{
			"message": helpers.FileError.Error(),
		})
		context.Abort()
		return
	}
	err = context.SaveUploadedFile(file, "csvFiles/"+file.Filename)
	categoryNames, err := helpers.ReadCsv("csvFiles/" + file.Filename)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	var existCategoryNames []string
	for _, categoryName := range categoryNames {
		err := c.CategoryService.GetCategoryByName(categoryName)
		if err != nil {
			existCategoryNames = append(existCategoryNames, categoryName)
			continue
		}
		c := category.Category{
			CategoryName: categoryName,
		}
		category.Create(&c)
	}

	if len(existCategoryNames) > 0 {
		context.JSON(http.StatusAlreadyReported, gin.H{
			"message":            "Category already exist others are created",
			"existCategoryNames": existCategoryNames,
		})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Category created successfully",
	})
}

func (c Controller) GetCategoryList(context *gin.Context) {
	categories := category.FindAll()

	var categoryList []string
	for _, ca := range categories {
		categoryList = append(categoryList, ca.CategoryName)
	}
	context.JSON(http.StatusOK, gin.H{
		"message":    "Categories listed successfully",
		"categories": categoryList,
	})
}
