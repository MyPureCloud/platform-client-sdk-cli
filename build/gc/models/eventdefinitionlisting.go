package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventdefinitionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventdefinitionlistingDud struct { 
    

}

// Eventdefinitionlisting
type Eventdefinitionlisting struct { 
    // Entities
    Entities []Eventdefinition `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Eventdefinitionlisting) String() string {
     o.Entities = []Eventdefinition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventdefinitionlisting) MarshalJSON() ([]byte, error) {
    type Alias Eventdefinitionlisting

    if EventdefinitionlistingMarshalled {
        return []byte("{}"), nil
    }
    EventdefinitionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Eventdefinition `json:"entities"`
        *Alias
    }{

        
        Entities: []Eventdefinition{{}},
        

        Alias: (*Alias)(u),
    })
}

