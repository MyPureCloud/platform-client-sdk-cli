package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentaggregatequeryresponseDud struct { 
    

}

// Agenticvirtualagentaggregatequeryresponse
type Agenticvirtualagentaggregatequeryresponse struct { 
    // Results
    Results []Agenticvirtualagentaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentaggregatequeryresponse) String() string {
     o.Results = []Agenticvirtualagentaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentaggregatequeryresponse

    if AgenticvirtualagentaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Agenticvirtualagentaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Agenticvirtualagentaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

