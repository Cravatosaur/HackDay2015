package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "aws"
    "encoding/json"
)

type Page struct {
    Title string
    Body  []byte
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World  %s!", r.URL.Path[1:])
}


func main() {

  http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", htmlHandler)
  http.HandleFunc("/js/", jsHandler)
  http.HandleFunc("/watch/",cloudwatchHandler)
	http.ListenAndServe(":8080", nil)
}

func cloudwatchHandler(w http.ResponseWriter, r *http.Request) {

  cw := new(aws.CloudWatcher)
  splitPath := strings.Split(r.URL.Path, "/")
  fmt.Println(splitPath)
  fmt.Println(len(splitPath))
  for i := 0 ; i < len(splitPath) ; i++ {
    fmt.Println(splitPath[i])
  }

  if ( len(splitPath) <= 2 || splitPath[2] == ""  ) {
    resp, _ := cw.ListMetrics()
    nameSpaces := make(map[string]int)
    for _,element := range resp.Metrics {
      nameSpaces[*element.Namespace] = nameSpaces[*element.Namespace] + 1
    }
    fmt.Println(resp)
    jsonResp, _ := json.Marshal(nameSpaces)
    fmt.Fprintf(w, "%s",  jsonResp)
  } else if ( len(splitPath) <= 3 || splitPath[3] == "" ){
    cw.MetricName = splitPath[2]
    resp, _ := cw.ListMetrics()
    fmt.Println(resp)

    dimensionNames := make(map[string]map[string]int)
    for _,element := range resp.Metrics {
      for _,dimension := range element.Dimensions {
        if ( dimensionNames[*dimension.Name] == nil ) {
          dimensionNames[*dimension.Name] = make(map[string]int)
        }
        dimensionNames[*dimension.Name][*dimension.Value] = dimensionNames[*dimension.Name][*dimension.Value] + 1
      }
    }
    jsonResp, _ := json.Marshal(dimensionNames)
    fmt.Fprintf(w, "%s",  jsonResp)


    // cw.NameSpace = "AWS/" + splitPath[2]
    // resp, _ := cw.ListMetrics()
    // fmt.Println(resp)
    //
    // dimensionNames := make(map[string]map[string]int)
    // for _,element := range resp.Metrics {
    //   for _,dimension := range element.Dimensions {
    //     if ( dimensionNames[*dimension.Name] == nil ) {
    //       dimensionNames[*dimension.Name] = make(map[string]int)
    //     }
    //     dimensionNames[*dimension.Name][*dimension.Value] = dimensionNames[*dimension.Name][*dimension.Value] + 1
    //   }
    // }
    // jsonResp, _ := json.Marshal(dimensionNames)
    // fmt.Fprintf(w, "%s",  jsonResp)
  } else if ( len(splitPath) <= 5 || splitPath[5] == "" ) {
    cw.NameSpace = "AWS/" + splitPath[2]
    cw.DimensionName = splitPath[3]
    cw.DimensionValue = splitPath[4]
    resp, _ := cw.ListMetrics()
    metrics := make(map[string]int)
    for _,element := range resp.Metrics {
      metrics[*element.MetricName] = metrics[*element.MetricName] + 1
    }
    jsonResp, _ := json.Marshal(metrics)
    fmt.Fprintf(w, "%s",  jsonResp)

  } else {
    cw.NameSpace = "AWS/" + splitPath[2]
    cw.DimensionName = "instanceid"//splitPath[3]
    cw.DimensionValue = "i-27747be1"//splitPath[4]
    cw.MetricName = splitPath[5]
    resp, _ := cw.FetchMetric()
    fmt.Println(resp.String())
    for _,element := range resp.Datapoints {
        fmt.Println(*element)
    }
    jsonResp, _ := json.Marshal(resp)
    fmt.Fprintf(w, "%s",  jsonResp)
  }
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    p, err := loadPage( "view/index.html" )
    if err != nil {
        fmt.Println("---")
        fmt.Println(err)
        fmt.Println("---")
        http.Redirect(w, r, "/view/404.html", http.StatusFound)
    } else {
      fmt.Fprintf(w, "%s",  p.Body)
    }
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    fmt.Println("Render HTML " + title )
    p, err := loadPage( "view/" + title )
    if err != nil {
        fmt.Println("---")
        fmt.Println(err)
        fmt.Println("---")
        http.Redirect(w, r, "/view/404.html", http.StatusFound)
    } else {
      fmt.Fprintf(w, "%s",  p.Body)
    }
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/js/"):]
    fmt.Println("Render JS " + title )
    p, err := loadPage( "js/" + title )
    if err != nil {
        fmt.Println("---")
        fmt.Println(err)
        fmt.Println("---")
        http.Redirect(w, r, "/view/404.html", http.StatusFound)
    } else {
      fmt.Fprintf(w, "%s",  p.Body)
    }
}

func loadPage(title string) (*Page, error) {
    filename := title
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("---")
        fmt.Println(err)
        fmt.Println("---")
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
