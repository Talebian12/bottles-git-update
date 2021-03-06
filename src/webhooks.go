package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
  path = "/webhooks"
)

type fn func()

func SetupWebhooks(action fn) {
  hook, _ := github.New(github.Options.Secret(os.Getenv("GH_SECRET")))
  
  http.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
    payload, err := hook.Parse(r, github.ReleaseEvent, github.PushEvent)
    if err != nil {
      if err == github.ErrEventNotFound {
        fmt.Println("Current event is not a push");
      }
    }
    
    switch payload.(type) {
      case github.ReleasePayload:
        release := payload.(github.ReleasePayload)
        fmt.Println("\n======WEBHOOK======\n")
        fmt.Printf("%+v\n", release)
      case github.PushPayload:
	//push := payload.(github.PushPayload)
        fmt.Println("\n======WEBHOOK======\n")
        action()
        //fmt.Printf("%+v\n", push)
    }

  })

  http.ListenAndServe(":3000", nil)
}
