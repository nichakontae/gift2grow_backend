package firebase

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

func Init() {
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		logrus.Fatal("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	App = app
}
