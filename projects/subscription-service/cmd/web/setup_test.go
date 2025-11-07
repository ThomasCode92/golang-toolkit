package main

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"subscription-service/data"

	"github.com/alexedwards/scs/v2"
)

var testApp Config

func TestMain(m *testing.M) {
	gob.Register(data.User{})

	// set up session
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	testApp = Config{
		Session:       session,
		DB:            nil,
		InfoLog:       log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Wait:          &sync.WaitGroup{},
		Models:        data.TestNew(nil),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	// create a dummy mailer
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerChanDone := make(chan bool)

	testApp.Mailer = Mail{
		Wait:       testApp.Wait,
		ErrChan:    errorChan,
		MailerChan: mailerChan,
		DoneChan:   mailerChanDone,
	}

	go func() {
		select {
		case <-testApp.Mailer.MailerChan:
		case <-testApp.Mailer.ErrChan:
		case <-testApp.Mailer.DoneChan:
			return
		}
	}()

	go func() {
		for {
			select {
			case err := <-testApp.ErrorChan:
				testApp.ErrorLog.Println(err)
			case <-testApp.ErrorChanDone:
				return
			}
		}
	}()

	os.Exit(m.Run())
}

func getCtx(req *http.Request) context.Context {
	ctx, err := testApp.Session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}

	return ctx
}
