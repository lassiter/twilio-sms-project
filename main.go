package main

import (
      "bufio"
      "encoding/csv"
      "encoding/json"
      "fmt"
      "io"
      "log"
      "os"
)

type person struct {
    name string
    countrycode int
    phonenumber int
    message string
}

func main() {

     accountSid := os.Getenv("ACCOUNT_ID")
     authToken := os.Getenv("AUTH_TOKEN")
     urlStr := "https://api.twilio.com/2010-04-01/Accounts/"   accountSid   "/Messages.json"
     

}