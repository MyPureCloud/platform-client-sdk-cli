package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenmonitoringsessionentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenmonitoringsessionentitylistingDud struct { 
    

}

// Screenmonitoringsessionentitylisting
type Screenmonitoringsessionentitylisting struct { 
    // Entities
    Entities []Screenmonitoringsession `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Screenmonitoringsessionentitylisting) String() string {
     o.Entities = []Screenmonitoringsession{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenmonitoringsessionentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Screenmonitoringsessionentitylisting

    if ScreenmonitoringsessionentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ScreenmonitoringsessionentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Screenmonitoringsession `json:"entities"`
        *Alias
    }{

        
        Entities: []Screenmonitoringsession{{}},
        

        Alias: (*Alias)(u),
    })
}

