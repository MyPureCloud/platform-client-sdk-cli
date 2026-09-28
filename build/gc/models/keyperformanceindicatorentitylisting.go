package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KeyperformanceindicatorentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KeyperformanceindicatorentitylistingDud struct { 
    

}

// Keyperformanceindicatorentitylisting
type Keyperformanceindicatorentitylisting struct { 
    // Entities
    Entities []Keyperformanceindicator `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Keyperformanceindicatorentitylisting) String() string {
     o.Entities = []Keyperformanceindicator{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Keyperformanceindicatorentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Keyperformanceindicatorentitylisting

    if KeyperformanceindicatorentitylistingMarshalled {
        return []byte("{}"), nil
    }
    KeyperformanceindicatorentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Keyperformanceindicator `json:"entities"`
        *Alias
    }{

        
        Entities: []Keyperformanceindicator{{}},
        

        Alias: (*Alias)(u),
    })
}

