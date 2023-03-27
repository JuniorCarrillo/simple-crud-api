package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/JuniorCarrillo/simple-crud-api/api/dtos"
	"github.com/JuniorCarrillo/simple-crud-api/api/models"
	"github.com/JuniorCarrillo/simple-crud-api/api/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func (db *DB) SaveCharacter(in *dtos.Character) (*models.Character, error) {
	dbColl := os.Getenv("MONGO_COLLECTION")

	coll := db.client.Database(dbColl).Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()
	characterToSave := dtos.CharacterToSave(in)
	characterToSave.Name = tools.ToCapitalize(characterToSave.Name)
	saved, err := coll.InsertOne(ctx, characterToSave)
	if err != nil {
		return nil, err
	}

	return dtos.CharacterToResponse(in, saved, now), nil
}

func (db *DB) Characters() ([]*models.Character, error) {
	dbColl := os.Getenv("MONGO_COLLECTION")
	coll := db.client.Database(dbColl).Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var characters []*models.Character
	for cur.Next(ctx) {
		var character *models.Character
		err = cur.Decode(&character)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return characters, nil
}

func (db *DB) CharactersSearch(search string, limit int64, page int64, url string) (*dtos.CharacterPaginationResponse, error) {
	dbColl := os.Getenv("MONGO_COLLECTION")
	if page < 1 {
		return nil, errors.New("use positive elements (greater than zero) for your pagination")
	}

	skip := page*limit - limit
	pagination := options.FindOptions{Limit: &limit, Skip: &skip}

	coll := db.client.Database(dbColl).Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	searchTerm := fmt.Sprintf(".*%s.*", search)
	filter := bson.D{{"name", primitive.Regex{Pattern: searchTerm}}}

	cur, err := coll.Find(ctx, filter, &pagination)
	if err != nil {
		return nil, err
	}

	var dataColl *mongo.Cursor
	dataColl, err = coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	totalColl := dataColl.RemainingBatchLength()

	var characters []*models.Character
	for cur.Next(ctx) {
		var character *models.Character
		err = cur.Decode(&character)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return dtos.CharacterToPagination(characters, int64(totalColl), page, limit, url, search), nil
}

func (db *DB) CharacterById(Id string) (*models.Character, error) {
	dbColl := os.Getenv("MONGO_COLLECTION")
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, err
	}

	coll := db.client.Database(dbColl).Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := coll.FindOne(ctx, bson.M{"_id": id})
	character := models.Character{}
	err = res.Decode(&character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (db *DB) CharacterUpdateById(Id string, up dtos.CharacterUpdate) (*models.Character, error) {
	nameCollection := os.Getenv("MONGO_COLLECTION")
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, err
	}

	coll := db.client.Database(nameCollection).Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": dtos.CharacterToUpdate(&up)}

	res := coll.FindOneAndUpdate(ctx, bson.M{"_id": id}, update)
	character := models.Character{}
	err = res.Decode(&character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (db *DB) CharacterDeleteById(Id string) (*models.Character, error) {
	nameCollection := os.Getenv("MONGO_COLLECTION")
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return nil, err
	}

	coll := db.client.Database(nameCollection).Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := coll.FindOneAndDelete(ctx, bson.M{"_id": id})
	character := models.Character{}
	err = res.Decode(&character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}
