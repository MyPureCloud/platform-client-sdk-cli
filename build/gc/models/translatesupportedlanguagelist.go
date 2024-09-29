package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranslatesupportedlanguagelistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranslatesupportedlanguagelistDud struct { 
    

}

// Translatesupportedlanguagelist
type Translatesupportedlanguagelist struct { 
    // Entities
    Entities []Translatesupportedlanguage `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Translatesupportedlanguagelist) String() string {
     o.Entities = []Translatesupportedlanguage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Translatesupportedlanguagelist) MarshalJSON() ([]byte, error) {
    type Alias Translatesupportedlanguagelist

    if TranslatesupportedlanguagelistMarshalled {
        return []byte("{}"), nil
    }
    TranslatesupportedlanguagelistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Translatesupportedlanguage `json:"entities"`
        *Alias
    }{

        
        Entities: []Translatesupportedlanguage{{}},
        

        Alias: (*Alias)(u),
    })
}

