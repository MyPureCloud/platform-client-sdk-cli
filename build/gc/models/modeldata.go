package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ModeldataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ModeldataDud struct { 
    


    

}

// Modeldata
type Modeldata struct { 
    // DownloadUrls - List of URLs to fetch the results of the forecast model. This field is populated only if session state is Complete
    DownloadUrls []string `json:"downloadUrls"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult []Planninggroupmodel `json:"downloadResult"`

}

// String returns a JSON representation of the model
func (o *Modeldata) String() string {
     o.DownloadUrls = []string{""} 
     o.DownloadResult = []Planninggroupmodel{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Modeldata) MarshalJSON() ([]byte, error) {
    type Alias Modeldata

    if ModeldataMarshalled {
        return []byte("{}"), nil
    }
    ModeldataMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrls []string `json:"downloadUrls"`
        
        DownloadResult []Planninggroupmodel `json:"downloadResult"`
        *Alias
    }{

        
        DownloadUrls: []string{""},
        


        
        DownloadResult: []Planninggroupmodel{{}},
        

        Alias: (*Alias)(u),
    })
}

