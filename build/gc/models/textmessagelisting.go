package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextmessagelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextmessagelistingDud struct { 
    

}

// Textmessagelisting
type Textmessagelisting struct { 
    // Entities
    Entities []Messagedata `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Textmessagelisting) String() string {
     o.Entities = []Messagedata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textmessagelisting) MarshalJSON() ([]byte, error) {
    type Alias Textmessagelisting

    if TextmessagelistingMarshalled {
        return []byte("{}"), nil
    }
    TextmessagelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Messagedata `json:"entities"`
        *Alias
    }{

        
        Entities: []Messagedata{{}},
        

        Alias: (*Alias)(u),
    })
}

