package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeoffplanlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeoffplanlistingDud struct { 
    


    

}

// Butimeoffplanlisting
type Butimeoffplanlisting struct { 
    // Entities
    Entities []Butimeoffplanresponse `json:"entities"`


    // DownloadUrl - URL from which to fetch results for requests with a large result set. If populated, the downloaded data will conform to the same schema as would normally be returned, excepting downloaded data will never itself contain a downloadUrl
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Butimeoffplanlisting) String() string {
     o.Entities = []Butimeoffplanresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeoffplanlisting) MarshalJSON() ([]byte, error) {
    type Alias Butimeoffplanlisting

    if ButimeoffplanlistingMarshalled {
        return []byte("{}"), nil
    }
    ButimeoffplanlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Butimeoffplanresponse `json:"entities"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        
        Entities: []Butimeoffplanresponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}

