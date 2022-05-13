package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitlistingDud struct { 
    

}

// Businessunitlisting
type Businessunitlisting struct { 
    // Entities
    Entities []Businessunitlistitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Businessunitlisting) String() string {
     o.Entities = []Businessunitlistitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitlisting) MarshalJSON() ([]byte, error) {
    type Alias Businessunitlisting

    if BusinessunitlistingMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Businessunitlistitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Businessunitlistitem{{}},
        

        Alias: (*Alias)(u),
    })
}

