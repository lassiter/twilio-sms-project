package main

import (
      "bufio"
      "encoding/csv"
      "encoding/json"
      "fmt"
      "io"
      "log"
      "os"
      "net/http"
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
   
   client := &http.Client{}
   
   req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
   
   req.SetBasicAuth(accountSid, authToken)
   
   req.Header.Add("Accept", "application/json")
   
   req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

   resp, _ := client.Do(req)
   
if (resp.StatusCode >= 200 && resp.StatusCode < 300) {

  var data map[string]interface{}
  
  decoder := json.NewDecoder(resp.Body)
  
  err := decoder.Decode(&data)
  if (err == nil) {
    fmt.Println(data["sid"])
  }
} else {
  fmt.Println(resp.Status);
}

}