package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceexplanationlistingbuqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceexplanationlistingbuqueryresponseDud struct { 
    

}

// Adherenceexplanationlistingbuqueryresponse
type Adherenceexplanationlistingbuqueryresponse struct { 
    // Entities
    Entities []Adherenceexplanationresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Adherenceexplanationlistingbuqueryresponse) String() string {
     o.Entities = []Adherenceexplanationresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceexplanationlistingbuqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Adherenceexplanationlistingbuqueryresponse

    if AdherenceexplanationlistingbuqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AdherenceexplanationlistingbuqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Adherenceexplanationresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Adherenceexplanationresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

