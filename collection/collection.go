package collection

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"shopping-service/driver"
	"shopping-service/models"
)

type Collection interface {
	GetCollection() (models.Result, error)
	GetCategories() (models.Result, error)
}

type collection struct {
	collection string
	executer driver.MongoConnector
}

func NewCollection(d driver.MongoConnector, c string) Collection {
	return &collection{executer:  d, collection: c}
}

func (c *collection) GetCollection() (models.Result, error) {
	filter := bson.D{{}}
	result, err := c.executer.Reader(filter, c.collection)

	if err != nil {
		return models.Result{}, fmt.Errorf("unable to get collections: %s", err)
	}

	return result.(models.Result), nil
}

func (c *collection) GetCategories() (models.Result, error) {
	filter := bson.D{{}}
	result, err := c.executer.Reader(filter, c.collection)

	if err != nil {
		return models.Result{}, fmt.Errorf("unable to get collections")
	}

	return result.(models.Result), nil
}