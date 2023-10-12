package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatmessageentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatmessageentitylistingDud struct { 
    

}

// Chatmessageentitylisting
type Chatmessageentitylisting struct { 
    // Entities
    Entities []Chatmessageresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Chatmessageentitylisting) String() string {
     o.Entities = []Chatmessageresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatmessageentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Chatmessageentitylisting

    if ChatmessageentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ChatmessageentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Chatmessageresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Chatmessageresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

