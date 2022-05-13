package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotsearchresponseentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotsearchresponseentitylistingDud struct { 
    

}

// Botsearchresponseentitylisting
type Botsearchresponseentitylisting struct { 
    // Entities
    Entities []Botsearchresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Botsearchresponseentitylisting) String() string {
     o.Entities = []Botsearchresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botsearchresponseentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Botsearchresponseentitylisting

    if BotsearchresponseentitylistingMarshalled {
        return []byte("{}"), nil
    }
    BotsearchresponseentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Botsearchresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Botsearchresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

