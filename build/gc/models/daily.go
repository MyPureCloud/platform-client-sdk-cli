package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DailyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DailyDud struct { 
    


    

}

// Daily
type Daily struct { 
    // DownloadUrl - Download URL to fetch the result of daily time series. This field is populated only if session state is Complete
    DownloadUrl string `json:"downloadUrl"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult []Timeseries `json:"downloadResult"`

}

// String returns a JSON representation of the model
func (o *Daily) String() string {
    
     o.DownloadResult = []Timeseries{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Daily) MarshalJSON() ([]byte, error) {
    type Alias Daily

    if DailyMarshalled {
        return []byte("{}"), nil
    }
    DailyMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadResult []Timeseries `json:"downloadResult"`
        *Alias
    }{

        


        
        DownloadResult: []Timeseries{{}},
        

        Alias: (*Alias)(u),
    })
}

