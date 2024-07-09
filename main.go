package main

import (
  "fmt"
  "strings"
  "regexp"
  "os"
  "path"
  "slices"
  "github.com/go-resty/resty/v2"
  "github.com/gocolly/colly/v2"
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
  PacerCaseId string `json:"pacer_case_id"`

}


func main () {

  // handle API info
  // can get unauthenticated of main docket info
  COURT_LISTENER_API := "https://www.courtlistener.com/api/rest/v3/dockets/"
  downloadDir := findUserFolder("Downloads")
  //fmt.Println(API_HEADER)

  // fmt.Println(os.Args)
  var case_spec, case_id string

  switch {
    case len(os.Args) > 1: {

      // fmt.Println("Reading case id from parameter")
      case_spec = os.Args[1]
    }
    default: {
      fmt.Print("Enter the case's courtlistener.com url: ")
      fmt.Scanln(&case_spec)
    }


  }

  fmt.Println("Attempting to find case")
  re := regexp.MustCompile(`\d+`)
  case_id = re.FindString(case_spec)

  fmt.Println("Searching CourtListener API ...")


  // Create REST client
  client := resty.New()

  caseData := ApiResult{}


  resp, err := client.R().
  EnableTrace().
  SetResult(&caseData).
  Get(COURT_LISTENER_API + case_id + "/")


  resp.Status()

  if err != nil {
    fmt.Println("Error, couldn't find the case.")
    os.Exit(0)
  } 

  if caseData.CaseName == "" {
    fmt.Println("Case not found")
    os.Exit(0)
  } 


  var caseYear string

  if caseData.DateFiled != "" {
    caseYear = caseDatate -output recap_downloader_universal recap_downloader_arm64 recap_downloader_amd64
    :!
    .DateFiled[2:4]
  } else {
    fmt.Println("Sorry, getting back invalid case data")
    os.Exit(1)
  }


  fmt.Println("Found", caseData.CaseName)
  fmt.Println("************** Downloading available PDFs/media from the Internet Archive ***************")

  c := colly.NewCollector(
  )


  c.OnRequest(func(r *colly.Request) {
    fmt.Print(".")
  })

  c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    e.Request.Visit(e.Attr("href"))
  })

  c.OnScraped(func(r *colly.Response) { 
    url := r.Request.URL.String()
    urlArray := strings.Split(url, "/")
    fileName := urlArray[len(urlArray) - 1]

    dotArray := strings.Split(url, ".")
    fileExtension := dotArray[len(dotArray) -1]
    junkFileTypes := []string{ "sqlite", "sql", "torrent" }

    if slices.Index(junkFileTypes, fileExtension) == -1 {
      r.Save(path.Join(downloadDir, fileName))
    }


  })


  c.Visit("https://ia802709.us.archive.org/" +
  caseYear +
  "/items/gov.uscourts." + 
  caseData.CourtId + "." + 
  string(caseData.PacerCaseId))

}





func findUserFolder(targetDirName string) string {
  // this function will be used to find the directory of name *name* within the user's
  // home directory on *nix and Windows. Supported directorys are "Downloads", "Documents", "Pictutres",
  // "Movies", etc.

  // depends only on the stdlib os module and some general knowledge about directories

  // throws an error if it fails
  userHomeDir, err := os.UserHomeDir()

  if err != nil {
    fmt.Println("error: ", err)
    os.Exit(1)
  } 

 return path.Join(userHomeDir, targetDirName)

}
