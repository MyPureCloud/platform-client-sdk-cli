package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentcopilotasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentcopilotasyncaggregatequeryresponseDud struct { 
    


    

}

// Agentcopilotasyncaggregatequeryresponse
type Agentcopilotasyncaggregatequeryresponse struct { 
    // Results
    Results []Agentcopilotaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Agentcopilotasyncaggregatequeryresponse) String() string {
     o.Results = []Agentcopilotaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentcopilotasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentcopilotasyncaggregatequeryresponse

    if AgentcopilotasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentcopilotasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Agentcopilotaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Agentcopilotaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

