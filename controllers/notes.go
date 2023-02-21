package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"netforemost/models"
	"netforemost/utils/responses"
)


func  GetNotes(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	query := r.URL.Query()
	order := query.Get("order")
	search := query.Get("search")

	// Set up filter and options for MongoDB query
	filter := bson.M{}
	options := options.Find()
	if search != "" {
		// Use regex to perform case-insensitive search on title and body fields
		filter["$or"] = bson.A{
			bson.M{"title": primitive.Regex{Pattern: search, Options: "i"}},
			bson.M{"body": primitive.Regex{Pattern: search, Options: "i"}},
		}
	}
	if order != "" {
		switch order {
		case "title":
			options.SetSort(bson.M{"title": 1})
		case "-title":
			options.SetSort(bson.M{"title": -1})
		case "created":
			options.SetSort(bson.M{"created": 1})
		case "-created":
			options.SetSort(bson.M{"created": -1})
		}
	}

	note := models.Note{}

	notes, err := note.FindNotes(s.DB, filter, options)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	responses.JSON(w, http.StatusAccepted, notes)
}

func CreateNote(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	note.Prepare()
	err = note.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	noteCreated, err := note.CreateNote(db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responses.JSON(w, http.StatusCreated, noteCreated)
}

func UpdateNote(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	var note models.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	return
	}
	
	idParam := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"title":   note.Title,
		"body":    note.Body,
		"created": time.Now(),
	}}

	result, err := note.UpdateNote(db, filter, update)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responses.JSON(w, http.StatusOK, result)
}

func DeleteNote(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	idParam := strings.TrimPrefix(r.URL.Path, "/notes/")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
	}

	result, err := models.DeleteNote(db, id)
	w.Header().Set("Content-Type", "application/json")
	responses.JSON(w, http.StatusOK, result)
}