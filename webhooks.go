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

func setup_webhooks() {
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
        fmt.Printf("%+v", release)
      case github.PushPayload:
        push := payload.(github.PushPayload)
        fmt.Printf("%+v", push)
    }

  })

  http.ListenAndServe(":3000", nil)
}
