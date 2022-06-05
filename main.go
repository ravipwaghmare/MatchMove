package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ravipwaghmare/MatchMove/database"
	"github.com/ravipwaghmare/MatchMove/server"

	"github.com/labstack/gommon/log"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	fmt.Println("Welcome to MatchMove App")

	logger := lumberjack.Logger{
		Filename:   "./Administrator.log",
		MaxAge:     10,
		MaxBackups: 10,
		Compress:   true, // disabled by default
		LocalTime:  true,
	}

	mw := io.MultiWriter(os.Stdout, &logger)
	log.SetOutput(mw)

	log.Print("CONNECT TO DATABASE")
	db, err := database.Connect()
	if err != nil {
		log.Panic("COULDN'T CONNECT TO DATABASE " + err.Error())
	}

	log.Print("START UP SERVER")
	apiServer, err := server.New(db, &logger)

	if err != nil {
		panic(fmt.Errorf("fatal error setting up server: %s", err))
	}

	apiServer.Logger.Fatal(apiServer.Start(":3000"))
}
