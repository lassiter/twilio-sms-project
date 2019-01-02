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
     
    msgData := url.Values{}

    msgData.Set("To","NUMBER_TO")

    msgData.Set("From","NUMBER_FROM")

    msgData.Set("Body",quotes[rand.Intn(len(quotes))])

   msgDataReader := *strings.NewReader(msgData.Encode())
   
   
     

}