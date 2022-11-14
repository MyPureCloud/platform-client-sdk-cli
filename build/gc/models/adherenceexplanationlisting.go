package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceexplanationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceexplanationlistingDud struct { 
    

}

// Adherenceexplanationlisting
type Adherenceexplanationlisting struct { 
    // Entities
    Entities []Adherenceexplanationresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Adherenceexplanationlisting) String() string {
     o.Entities = []Adherenceexplanationresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceexplanationlisting) MarshalJSON() ([]byte, error) {
    type Alias Adherenceexplanationlisting

    if AdherenceexplanationlistingMarshalled {
        return []byte("{}"), nil
    }
    AdherenceexplanationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Adherenceexplanationresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Adherenceexplanationresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

