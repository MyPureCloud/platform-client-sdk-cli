package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitactivitycodelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitactivitycodelistingDud struct { 
    


    

}

// Businessunitactivitycodelisting
type Businessunitactivitycodelisting struct { 
    // Entities
    Entities []Businessunitactivitycode `json:"entities"`


    // DownloadUrl - URL from which to fetch results for requests with a large result set. If populated, the downloaded data will conform to the same schema as would normally be returned, excepting downloaded data will never itself contain a downloadUrl
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Businessunitactivitycodelisting) String() string {
     o.Entities = []Businessunitactivitycode{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitactivitycodelisting) MarshalJSON() ([]byte, error) {
    type Alias Businessunitactivitycodelisting

    if BusinessunitactivitycodelistingMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitactivitycodelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Businessunitactivitycode `json:"entities"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        
        Entities: []Businessunitactivitycode{{}},
        


        

        Alias: (*Alias)(u),
    })
}

