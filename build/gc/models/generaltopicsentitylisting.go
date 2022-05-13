package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeneraltopicsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeneraltopicsentitylistingDud struct { 
    

}

// Generaltopicsentitylisting
type Generaltopicsentitylisting struct { 
    // Entities
    Entities []Generaltopic `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Generaltopicsentitylisting) String() string {
     o.Entities = []Generaltopic{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generaltopicsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Generaltopicsentitylisting

    if GeneraltopicsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    GeneraltopicsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Generaltopic `json:"entities"`
        *Alias
    }{

        
        Entities: []Generaltopic{{}},
        

        Alias: (*Alias)(u),
    })
}

