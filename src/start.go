package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)


func executeActivity() {
   dir, err := os.Getwd()
   out, err := exec.Command("/bin/bash", dir + "/update_ver.sh").Output()

   if err != nil {
      log.Fatal(err)
   } else {
      fmt.Println("Pushed to repository, updated files")
   }
   fmt.Printf("%s\n", out)
}

func main() {
   SetupWebhooks(executeActivity)  
}
