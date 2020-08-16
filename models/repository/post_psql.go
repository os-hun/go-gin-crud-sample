package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/os-hun/go-gin-crud-sample/db"
	"github.com/os-hun/go-gin-crud-sample/models"
	"github.com/os-hun/go-gin-crud-sample/api"
)

type PostRepository struct{}

type Post api.Post

// GetAll is get all Post
func (_ PostRepository) GetAll() ([]Post, error) {
    db := db.GetDB()
    var p []Post

    if err := db.Find(&p).Error; err != nil {
        return nil, err
    }

    return p, nil
}

// CreateModel is create Post model
func (_ PostRepository) CreateModel(p *models.Post) (*models.Post, error) {
    db := db.GetDB()
    if err := db.Create(&p).Error; err != nil {
        return p, err
    }
    return p, nil
}

// GetByID is get a Post by ID
func (_ PostRepository) GetByID(id int) (models.Post, error) {
    db := db.GetDB()
    var p models.Post
    if err := db.Where("id = ?", id).First(&p).Error; err != nil {
        return p, err
    }
    return p, nil
}

// UpdateByID is update a Post
func (_ PostRepository) UpdateByID(id int, c *gin.Context) (api.Post, error) {
    db := db.GetDB()
    var p api.Post
    if err := db.Where("id = ?", id).First(&p).Error; err != nil {
        return p, err
    }
    userID := p.UserID
    if err := c.BindJSON(&p); err != nil {
        return p, err
    }
    fmt.Printf("%v", p)
    p.ID = uint(id)
    p.UserID = userID
    db.Save(&p)

    return p, nil
}

// DeleteByID is delete a Post by ID
func (_ PostRepository) DeleteByID(id int) error {
    db := db.GetDB()
    var p Post

    if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
        return err
    }

    return nil
}
