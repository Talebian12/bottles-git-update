package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)


func executeActivity() {
   dir, err := os.Getwd()
   cmd := exec.Command("bash " + dir + "/update_ver.sh")

   err = cmd.Run()

   if err != nil {
      log.Fatal(err)
   } else {
      fmt.Println("Pushed to repository, updated files")
   }
}

func main() {
   SetupWebhooks(executeActivity)  
}
