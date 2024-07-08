package main

import (
  "fmt"
  //  "net/http"
  "os"
  //  "io"
  "regexp"
  "github.com/go-resty/resty/v2"
)

type ApiResult struct {
  ResourceUri string `json:"resource_uri"`
  Id int `json:"id"`
  Court string `json:"court"`
  CourtId string `json:"court_id"`
  AbosluteUrl string `json:"absolute_url"`
  DateFiled string `json:"date_filed"`
  DateModified string `json:"date_modified"`
  AssignedTo string `json:"assigned_to_string"`
  CaseName string `json:"case_name"`
  DateLastFiling string `json:"date_last_filing"`
}


func main () {

  // handle API info
  // can get unauthenticated of main docket info
  COURT_LISTENER_API := "https://www.courtlistener.com/api/rest/v3/dockets/"

  //fmt.Println(API_HEADER)

  fmt.Println(os.Args)
  var case_spec, case_id string

  switch {
    case len(os.Args) > 1: {

      fmt.Println("Reading case id from parameter")
      case_spec = os.Args[1]
    }
    default: {
      fmt.Print("Enter case id or url: ")
      fmt.Scanln(&case_spec)
    }


  }

  fmt.Println("Attempting to find case")
  re := regexp.MustCompile(`\d+`)
  case_id = re.FindString(case_spec)

  fmt.Println("Searching CourtListener API for", case_id)


  // Create REST client
  client := resty.New()

  caseData := ApiResult{}


  resp, err := client.R().
  EnableTrace().
  SetResult(&caseData).
  Get(COURT_LISTENER_API + case_id + "/")

  
  if err != nil {
    fmt.Println("error: ", err)
    os.Exit(1)
  } 

  fmt.Println(caseData)
}


