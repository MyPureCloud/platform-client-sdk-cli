package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserstaffinggrouplistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserstaffinggrouplistingDud struct { 
    


    

}

// Userstaffinggrouplisting
type Userstaffinggrouplisting struct { 
    // Entities
    Entities []Userstaffinggroupresponse `json:"entities"`


    // DownloadUrl - URL from which to fetch results for requests with a large result set. If populated, the downloaded data will conform to the same schema as would normally be returned, excepting downloaded data will never itself contain a downloadUrl
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Userstaffinggrouplisting) String() string {
     o.Entities = []Userstaffinggroupresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userstaffinggrouplisting) MarshalJSON() ([]byte, error) {
    type Alias Userstaffinggrouplisting

    if UserstaffinggrouplistingMarshalled {
        return []byte("{}"), nil
    }
    UserstaffinggrouplistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userstaffinggroupresponse `json:"entities"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        
        Entities: []Userstaffinggroupresponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}

