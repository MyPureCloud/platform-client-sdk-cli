package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuarterhourlyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuarterhourlyDud struct { 
    


    

}

// Quarterhourly
type Quarterhourly struct { 
    // DownloadUrls - List of download URLs to fetch the result of quarter hour time series. This field is populated only if session state is Complete
    DownloadUrls []string `json:"downloadUrls"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult []Timeseries `json:"downloadResult"`

}

// String returns a JSON representation of the model
func (o *Quarterhourly) String() string {
     o.DownloadUrls = []string{""} 
     o.DownloadResult = []Timeseries{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Quarterhourly) MarshalJSON() ([]byte, error) {
    type Alias Quarterhourly

    if QuarterhourlyMarshalled {
        return []byte("{}"), nil
    }
    QuarterhourlyMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrls []string `json:"downloadUrls"`
        
        DownloadResult []Timeseries `json:"downloadResult"`
        *Alias
    }{

        
        DownloadUrls: []string{""},
        


        
        DownloadResult: []Timeseries{{}},
        

        Alias: (*Alias)(u),
    })
}

