package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MaskingrulelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MaskingrulelistingDud struct { 
    

}

// Maskingrulelisting
type Maskingrulelisting struct { 
    // Entities
    Entities []Maskingrule `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Maskingrulelisting) String() string {
     o.Entities = []Maskingrule{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Maskingrulelisting) MarshalJSON() ([]byte, error) {
    type Alias Maskingrulelisting

    if MaskingrulelistingMarshalled {
        return []byte("{}"), nil
    }
    MaskingrulelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Maskingrule `json:"entities"`
        *Alias
    }{

        
        Entities: []Maskingrule{{}},
        

        Alias: (*Alias)(u),
    })
}

