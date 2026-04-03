package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluateshifttradelistjobresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluateshifttradelistjobresponseitemDud struct { 
    

}

// Evaluateshifttradelistjobresponseitem
type Evaluateshifttradelistjobresponseitem struct { 
    // Entities
    Entities []Evaluatedshifttraderesponseitem `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Evaluateshifttradelistjobresponseitem) String() string {
     o.Entities = []Evaluatedshifttraderesponseitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluateshifttradelistjobresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Evaluateshifttradelistjobresponseitem

    if EvaluateshifttradelistjobresponseitemMarshalled {
        return []byte("{}"), nil
    }
    EvaluateshifttradelistjobresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Evaluatedshifttraderesponseitem `json:"entities"`
        *Alias
    }{

        
        Entities: []Evaluatedshifttraderesponseitem{{}},
        

        Alias: (*Alias)(u),
    })
}

