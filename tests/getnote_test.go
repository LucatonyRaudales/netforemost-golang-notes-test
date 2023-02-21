package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
	"fmt"
	"testing"
	"netforemost/models"
	"netforemost/controllers"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)
func TestCreateAndGetNotes(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://hindra:vUxNd2AF3pTWHkgb@hindra-auth.b8lhl8z.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	title, body := getRandomTitleAndBody()
	note1 := models.Note{Title: title, Body: body}

	_, err = note1.CreateNote(client)
	if err != nil {
	    t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/notes?search=" + title, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	controllers.GetNotes(client, rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v want %v", rr.Code, http.StatusOK)
	}

	var notes []models.Note
	err = json.Unmarshal(rr.Body.Bytes(), &notes)
	if err != nil {
		t.Fatal(err)
	}
	if len(notes) < 1 {
		t.Errorf("unexpected number of notes: got %v but should have at least %v", len(notes), 1)
	}
	if notes[0].Title != title {
		t.Errorf("unexpected note titles: got %v but iwant '%v'", notes[0].Title, title)
	}
	if notes[0].Body != body {
		t.Errorf("unexpected note body: got %v but iwant '%v''", notes[0].Body, body)
	}
}

func getRandomTitleAndBody() (string, string) {
    t := time.Now()
    title := fmt.Sprintf("Title Test Note %d%d%d", t.Hour(), t.Minute(), t.Second())
    body := fmt.Sprintf("Body Test, This is a note created at %d:%d:%d", t.Hour(), t.Minute(), t.Second())
    return title, body
}