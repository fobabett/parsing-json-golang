package main

import "net/http"
import "io/ioutil"
import "log"
import "encoding/json"

type reddit struct {
  Kind string `json:"kind"`
  Data ytData `json:"data"`
}

type ytData struct {
  Modhash string `json:"modhash"`
  Children []ytChildren `json:"children"`
}

type ytChildren struct {
  Kind string `json:"kind"`
}

func main() {
  resp, err := http.Get("http://www.reddit.com/r/aww.json")
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatal(err)
  }

  var y reddit
  err = json.Unmarshal(body, &y)

  if err != nil {
    log.Fatal(err)
  }

  log.Printf("%+v", y)
}