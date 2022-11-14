package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceexplanationlistingagentqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceexplanationlistingagentqueryresponseDud struct { 
    

}

// Adherenceexplanationlistingagentqueryresponse
type Adherenceexplanationlistingagentqueryresponse struct { 
    // Entities
    Entities []Adherenceexplanationresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Adherenceexplanationlistingagentqueryresponse) String() string {
     o.Entities = []Adherenceexplanationresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceexplanationlistingagentqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Adherenceexplanationlistingagentqueryresponse

    if AdherenceexplanationlistingagentqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    AdherenceexplanationlistingagentqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Adherenceexplanationresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Adherenceexplanationresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

