package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutlierdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutlierdataDud struct { 
    


    

}

// Outlierdata
type Outlierdata struct { 
    // DownloadUrls - List of URLs to fetch the result of the outliers data. This field is populated only if session state is 'Complete' and outliers are present
    DownloadUrls []string `json:"downloadUrls"`


    // DownloadResultTemplate - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResultTemplate Outlierresulttemplate `json:"downloadResultTemplate"`

}

// String returns a JSON representation of the model
func (o *Outlierdata) String() string {
     o.DownloadUrls = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outlierdata) MarshalJSON() ([]byte, error) {
    type Alias Outlierdata

    if OutlierdataMarshalled {
        return []byte("{}"), nil
    }
    OutlierdataMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrls []string `json:"downloadUrls"`
        
        DownloadResultTemplate Outlierresulttemplate `json:"downloadResultTemplate"`
        *Alias
    }{

        
        DownloadUrls: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

