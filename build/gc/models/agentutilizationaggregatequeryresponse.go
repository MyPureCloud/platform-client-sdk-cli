package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentutilizationaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentutilizationaggregatequeryresponseDud struct { 
    

}

// Agentutilizationaggregatequeryresponse
type Agentutilizationaggregatequeryresponse struct { 
    // Results
    Results []Agentutilizationaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Agentutilizationaggregatequeryresponse) String() string {
     o.Results = []Agentutilizationaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentutilizationaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentutilizationaggregatequeryresponse

    if AgentutilizationaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentutilizationaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Agentutilizationaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Agentutilizationaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

