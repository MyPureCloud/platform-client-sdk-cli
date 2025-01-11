package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeeklyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeeklyDud struct { 
    


    

}

// Weekly
type Weekly struct { 
    // DownloadUrl - Download URL to fetch the result of weekly time series. This field is populated only if session state is Complete
    DownloadUrl string `json:"downloadUrl"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult []Timeseries `json:"downloadResult"`

}

// String returns a JSON representation of the model
func (o *Weekly) String() string {
    
     o.DownloadResult = []Timeseries{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekly) MarshalJSON() ([]byte, error) {
    type Alias Weekly

    if WeeklyMarshalled {
        return []byte("{}"), nil
    }
    WeeklyMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadResult []Timeseries `json:"downloadResult"`
        *Alias
    }{

        


        
        DownloadResult: []Timeseries{{}},
        

        Alias: (*Alias)(u),
    })
}

