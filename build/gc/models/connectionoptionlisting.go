package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectionoptionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectionoptionlistingDud struct { 
    

}

// Connectionoptionlisting
type Connectionoptionlisting struct { 
    // Entities
    Entities []Connectionoption `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Connectionoptionlisting) String() string {
     o.Entities = []Connectionoption{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectionoptionlisting) MarshalJSON() ([]byte, error) {
    type Alias Connectionoptionlisting

    if ConnectionoptionlistingMarshalled {
        return []byte("{}"), nil
    }
    ConnectionoptionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Connectionoption `json:"entities"`
        *Alias
    }{

        
        Entities: []Connectionoption{{}},
        

        Alias: (*Alias)(u),
    })
}

