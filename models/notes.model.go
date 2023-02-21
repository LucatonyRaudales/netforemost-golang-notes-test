package models

import (
	"context"
	"errors"
	"html"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Note struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Body    string             `bson:"body"`
	Created time.Time          `bson:"created"`
}

type DeleteStruct struct {
	ID primitive.ObjectID `bson:"note_id,omitempty"`
}


const (
	collectionName 	= "Notes"
	databaseName 	= "netforemost"
)

func (data *Note) Prepare(){
	data.Title = html.EscapeString(strings.TrimSpace((data.Title)))
	data.Body = html.EscapeString(strings.TrimSpace((data.Body)))
	data.Created = time.Now()
}

func (data *Note) Validate() error {
	if data.Title == "" {
		return errors.New("title is required")
	}
	if data.Body == "" {
		return errors.New("body is required")
	}
	return nil
}

func (note *Note) CreateNote(db *mongo.Client) (*mongo.InsertOneResult, error) {
	var err error
	collection := db.Database(databaseName).Collection(collectionName)
	result, err := collection.InsertOne(context.Background(), note)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (note *Note) FindNotes(db *mongo.Client, filter primitive.M, options *options.FindOptions) (*[]Note, error) {
	var err error
	collection := db.Database(databaseName).Collection(collectionName)
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return &[]Note{}, err
	}
	var notes []Note
	for cursor.Next(context.Background()) {
		var note Note
		err := cursor.Decode(&note)
		if err != nil {
			return &[]Note{}, err
		}
		notes = append(notes, note)
	}
	return &notes, nil
}

func (note *Note) UpdateNote(db *mongo.Client, filter, update primitive.M)(*mongo.UpdateResult, error){
	var err error

	collection := db.Database(databaseName).Collection(collectionName)
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteNote(db *mongo.Client, noteID primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": noteID}
	collection := db.Database(databaseName).Collection(collectionName)
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}