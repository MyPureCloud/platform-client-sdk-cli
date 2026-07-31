package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisionmetricsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisionmetricsresponseDud struct { 
    

}

// Decisionmetricsresponse
type Decisionmetricsresponse struct { 
    // Entities
    Entities []Decisionmetricsdata `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Decisionmetricsresponse) String() string {
     o.Entities = []Decisionmetricsdata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisionmetricsresponse) MarshalJSON() ([]byte, error) {
    type Alias Decisionmetricsresponse

    if DecisionmetricsresponseMarshalled {
        return []byte("{}"), nil
    }
    DecisionmetricsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Decisionmetricsdata `json:"entities"`
        *Alias
    }{

        
        Entities: []Decisionmetricsdata{{}},
        

        Alias: (*Alias)(u),
    })
}

