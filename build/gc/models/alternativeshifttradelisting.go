package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshifttradelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshifttradelistingDud struct { 
    

}

// Alternativeshifttradelisting
type Alternativeshifttradelisting struct { 
    // Entities
    Entities []Alternativeshifttraderesponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Alternativeshifttradelisting) String() string {
     o.Entities = []Alternativeshifttraderesponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshifttradelisting) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshifttradelisting

    if AlternativeshifttradelistingMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshifttradelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Alternativeshifttraderesponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Alternativeshifttraderesponse{{}},
        

        Alias: (*Alias)(u),
    })
}

