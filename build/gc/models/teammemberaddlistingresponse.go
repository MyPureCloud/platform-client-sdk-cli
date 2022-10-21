package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeammemberaddlistingresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeammemberaddlistingresponseDud struct { 
    


    Failures []Teamaddmemberfailure `json:"failures"`

}

// Teammemberaddlistingresponse
type Teammemberaddlistingresponse struct { 
    // Entities
    Entities []Userreference `json:"entities"`


    

}

// String returns a JSON representation of the model
func (o *Teammemberaddlistingresponse) String() string {
     o.Entities = []Userreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teammemberaddlistingresponse) MarshalJSON() ([]byte, error) {
    type Alias Teammemberaddlistingresponse

    if TeammemberaddlistingresponseMarshalled {
        return []byte("{}"), nil
    }
    TeammemberaddlistingresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userreference `json:"entities"`
        *Alias
    }{

        
        Entities: []Userreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

