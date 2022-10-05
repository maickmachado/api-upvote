package database

import (
	"github.com/maickmachado/upvote-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetCryptoInfoDataBase(s string) ([]*models.CryptoDataBase, error) {
	filter := bson.D{
		primitive.E{Key: "name", Value: s},
	}

	return FilterTasks(filter)
}

func FilterTasks(filter interface{}) ([]*models.CryptoDataBase, error) {
	// A slice of tasks for storing the decoded documents
	var tasks []*models.CryptoDataBase

	cur, err := Collection.Find(Ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(Ctx) {
		var t models.CryptoDataBase
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	// once exhausted, close the cursor
	err = cur.Close(Ctx)
	if err != nil {
		log.Println(err)
	}

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func Upvote(text string) {

	crypto, err := GetCryptoInfoDataBase(text)
	if err != nil {
		log.Println(err)
	}

	var name string
	if len(crypto) == 0 {
		name = ""
	} else {
		name = crypto[0].Name
	}

	if name == text {
		filter := bson.D{primitive.E{Key: "name", Value: text}}

		update := bson.D{primitive.E{Key: "$inc", Value: bson.D{
			primitive.E{Key: "votes", Value: 1},
		}}}

		t := &models.CryptoDataBase{}
		err = Collection.FindOneAndUpdate(Ctx, filter, update).Decode(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		task := &models.CryptoDataBase{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      text,
			Upvote:    1,
		}

		_, err := Collection.InsertOne(Ctx, task)
		if err != nil {
			log.Println(err)
		}
	}
}

func CheckIfExist(responseObject models.Response, detailResponseObject models.CryptoData) (*models.CryptoData, bool) {
	for _, value := range responseObject.CryptoData {
		if detailResponseObject.Slug == value.Slug {
			detailResponseObject = value
			return &detailResponseObject, true
		}
	}
	return nil, false
}

func OrderByVotes() ([]*models.CryptoDataBase, error) {

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"votes", -1}})

	var tasks []*models.CryptoDataBase

	cur, err := Collection.Find(Ctx, filter, opts)
	if err != nil {
		return tasks, err
	}

	for cur.Next(Ctx) {
		var t models.CryptoDataBase
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	// once exhausted, close the cursor
	err = cur.Close(Ctx)
	if err != nil {
		log.Println(err)
	}

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, err
}
