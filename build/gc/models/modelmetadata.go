package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ModelmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ModelmetadataDud struct { 
    


    

}

// Modelmetadata
type Modelmetadata struct { 
    // DownloadUrls - List of URLs to fetch the result of the forecast metadata. This field is populated only if session state is Complete
    DownloadUrls []string `json:"downloadUrls"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult Modelmetadataresult `json:"downloadResult"`

}

// String returns a JSON representation of the model
func (o *Modelmetadata) String() string {
     o.DownloadUrls = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Modelmetadata) MarshalJSON() ([]byte, error) {
    type Alias Modelmetadata

    if ModelmetadataMarshalled {
        return []byte("{}"), nil
    }
    ModelmetadataMarshalled = true

    return json.Marshal(&struct {
        
        DownloadUrls []string `json:"downloadUrls"`
        
        DownloadResult Modelmetadataresult `json:"downloadResult"`
        *Alias
    }{

        
        DownloadUrls: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

