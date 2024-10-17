package recipes

import (
	"encoding/json"
	"log/slog"

	"github.com/judewood/bakery/internal/s3client"
)

// Ingredient is a food ingredient for a product
type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// Recipe is the ingredients and instructions for creating a product
type Recipe struct {
	Name        string       `json:"name"`
	ID          string       `json:"id"`
	Ingredients []Ingredient `json:"ingredients"`
	BakeTime    int          `json:"bakeTime"`
}

// RecipeStorer contains CRUD methods for recipes
type RecipeStorer interface {
	GetRecipeFromS3(id string) (Recipe, error)
}

// RecipeStore implements crud operations on recipes
type RecipeStore struct {
	url string
	s3  s3client.S3Communicator
}

func New(url string, s3 s3client.S3Communicator) *RecipeStore {
	slog.Info("S3 recipe folder", "folder", url)
	return &RecipeStore{
		url: url,
		s3:  s3,
	}
}

func (r *RecipeStore) GetRecipeFromS3(id string) (Recipe, error) {
	url := r.url + id + ".json"
	slog.Debug("Getting S3 object from url", "url", url, "id", id)
	response, err := r.s3.GetDataFromS3(url)
	recipe := Recipe{}
	if err != nil {
		return Recipe{}, err
	}
	err = json.Unmarshal(response, &recipe)
	if err != nil {
		slog.Warn("failed to deserialise recipe form S3", "error", err)
		return Recipe{}, err
	}
	slog.Debug("received recipe from S3", "recipe", recipe)
	return recipe, nil
}
