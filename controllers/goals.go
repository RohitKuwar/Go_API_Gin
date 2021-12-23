package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/RohitKuwar/go_api_gin/config"
	"github.com/RohitKuwar/go_api_gin/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var Config, _ = config.LoadConfig(".")
var firestoreCredentialsLocation = Config.FirestoreCred

func GetGoals(c *gin.Context) {
	fmt.Println("ProjectId:", Config.ProjectId)

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	var newGoals []models.Goal
	iter := client.Collection("goals").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		fmt.Print(doc.Data())

		var tempGoals models.Goal
		if err := doc.DataTo(&tempGoals); err != nil {
			break
		}
		newGoals = append(newGoals, tempGoals)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"goals":   newGoals,
		"message": "Goals returned successfully!",
	})
}

func GetGoal(c *gin.Context) {
	// get parameter value
	paramID := c.Params.ByName("id")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	dsnap, err := client.Collection("goals").Doc(paramID).Get(ctx)
	if err != nil {
		fmt.Print(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Goal not found",
		})
		return
	}
	m := dsnap.Data()
	// fmt.Printf("Document data: %#v\n", m)

	// if m == nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{
	// 		"message": "doc does not exist",
	// 	})
	// }

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"Goal":    m,
		"message": "Goal with id " + paramID + " returned successfully!",
	})

}

func CreateGoal(c *gin.Context) {
	type Request struct {
		Id              string `json:"id"`
		Title           string `json:"title"`
		Status          string `json:"status"`
		AssignedTo      string `json:"assignedTo"`
		AssignedBy      string `json:"assignedBy"`
		AssignedOn      string `json:"assignedOn"`
		CompletionAward string `json:"completionAward"`
	}

	var body Request

	err := c.Bind(&body)

	// if error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// create a goal variable
	goal := &models.Goal{
		Title:           body.Title,
		Status:          body.Status,
		AssignedTo:      body.AssignedTo,
		AssignedBy:      body.AssignedBy,
		AssignedOn:      body.AssignedOn,
		CompletionAward: body.CompletionAward,
	}

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	ref := client.Collection("goals").NewDoc()

	goal.Id = ref.ID

	_, err = ref.Set(ctx, map[string]interface{}{
		"id":              goal.Id,
		"title":           goal.Title,
		"status":          goal.Status,
		"assignedTo":      goal.AssignedTo,
		"assignedBy":      goal.AssignedBy,
		"assignedOn":      goal.AssignedOn,
		"completionAward": goal.CompletionAward,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "data added successfully",
		"goal":    goal,
	})
}

func UpdateGoal(c *gin.Context) {
	type request struct {
		Title           string `json:"title"`
		Status          string `json:"status"`
		AssignedOn      string `json:"assignedOn"`
		CompletionAward string `json:"completionAward"`
	}
	var body request

	err := c.Bind(&body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot parse JSON",
		})
	}

	paramID := c.Params.ByName("id")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	fmt.Print(paramID, body.Status, body.Title)

	_, err = client.Collection("goals").Doc(paramID).Set(ctx, map[string]interface{}{
		"title":           body.Title,
		"status":          body.Status,
		"assignedOn":      body.AssignedOn,
		"completionAward": body.CompletionAward,
	}, firestore.MergeAll)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "data updated successfully",
	})
}

func DeleteGoal(c *gin.Context) {
	// get param
	paramID := c.Params.ByName("id")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	dsnap, err := client.Collection("goals").Doc(paramID).Get(ctx)
	if err != nil {
		fmt.Print(err)
		// if goal not found
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Goal with id " + paramID + " Not found",
		})
	}

	// Test Print line
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)

	_, err = client.Collection("goals").Doc(paramID).Delete(ctx)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Goal with id " + paramID + " Deleted Successfully",
	})
}
