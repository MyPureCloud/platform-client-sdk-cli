package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentcopilotaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentcopilotaggregatequeryresponseDud struct { 
    

}

// Agentcopilotaggregatequeryresponse
type Agentcopilotaggregatequeryresponse struct { 
    // Results
    Results []Agentcopilotaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Agentcopilotaggregatequeryresponse) String() string {
     o.Results = []Agentcopilotaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentcopilotaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentcopilotaggregatequeryresponse

    if AgentcopilotaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentcopilotaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Agentcopilotaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Agentcopilotaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

