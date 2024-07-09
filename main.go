package main

import (
  "fmt"
  "strings"
  "os"
  "io/fs"
  "regexp"
  "github.com/go-resty/resty/v2"
  "github.com/gocolly/colly/v2"
  "slices"
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


  resp.Status()

  if err != nil {
    fmt.Println("error: ", err)
    os.Exit(1)
  } 

  fmt.Println(resp.Result())


  caseYear := caseData.DateFiled[2:4]


  fmt.Println("************** download PDFs from Internet Archive ***************")

  c := colly.NewCollector()

  c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting", r.URL)
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

    fmt.Println(fileName)
    fmt.Println(fileExtension)

    r.Save(fileName)

  })


  c.Visit("https://ia802709.us.archive.org/" +
  caseYear +
  "/items/gov.uscourts." + 
  caseData.CourtId + "." + 
  string(caseData.PacerCaseId))

}





func findUserFolder(targetDirName string) {
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

  userDirsStructs, err := fs.ReadDir(userHomeDir)

  var targetDirIndex int

  targetDirNameSingular := strings.TrimRight(targetDirName, 1)
  targetDirNamePlural := targetDirName + "s"

  userDirs := map(userDirsStructs, func (dir) {dir.Name})

  switch { 
  case slices.Index(userDirs, targetDirName) > -1:
    targetDirIndex = slices.Index(userDirs, targetDirName) 

  case slices.Index(userDirs, strings.ToLower(targetDirName)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.ToLower(targetDirName)) 

  case slices.Index(userDirs, strings.ToUpper(targetDirName)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.ToUpper(targetDirName)) 

    // using deprecated call to Title. Hopefully the experimental text package's cases method will 
    // help
  case slices.Index(userDirs, strings.Title(targetDirName)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.Title(targetDirName)) 

    // with targetDirNamePlural 
  case slices.Index(userDirs, targetDirNamePlural) > -1:
    targetDirIndex = slices.Index(userDirs, targetDirNamePlural) 

  case slices.Index(userDirs, strings.ToLower(targetDirNamePlural)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.ToLower(targetDirNamePlural)) 

  case slices.Index(userDirs, strings.ToUpper(targetDirNamePlural)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.ToUpper(targetDirNamePlural)) 

    // using deprecated call to Title. Hopefully the experimental text package's cases method will 
    // help
  case slices.Index(userDirs, strings.Title(targetDirNamePlural)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.Title(targetDirNamePlural))

    // with targetDirNameSingular 
  case slices.Index(userDirs, targetDirNameSingular) > -1:
    targetDirIndex = slices.Index(userDirs, targetDirNameSingular) 

  case slices.Index(userDirs, strings.ToLower(targetDirNameSinglular)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.ToLower(targetDirNameSingular)) 

  case slices.Index(userDirs, strings.ToUpper(targetDirNameSingular)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.ToUpper(targetDirNameSingular)) 

    // using deprecated call to Title. Hopefully the experimental text package's cases method will 
    // help
  case slices.Index(userDirs, strings.Title(targetDirNameSingular)) > -1:
    targetDirIndex = slices.Index(userDirs, strings.Title(targetDirNameSingular)) 
  default:
    targetDirIndex = -1
  }

  if targetDirIndex > -1 {
    return path.join(userHomeDir, userDirsStructs[targetDirIndex].Name)
  }


 return userHomeDir

}

func findUserFolderWithCallback(name string, callback string) {
  return callback
  // eventually add a WithCallback function that fires a callback in lieu of failure, easily used to trigger 
  // user input of some kind or to failover to a temporary directory.
}
