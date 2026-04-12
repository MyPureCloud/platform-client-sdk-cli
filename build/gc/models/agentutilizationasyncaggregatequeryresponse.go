package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentutilizationasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentutilizationasyncaggregatequeryresponseDud struct { 
    


    

}

// Agentutilizationasyncaggregatequeryresponse
type Agentutilizationasyncaggregatequeryresponse struct { 
    // Results
    Results []Agentutilizationaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Agentutilizationasyncaggregatequeryresponse) String() string {
     o.Results = []Agentutilizationaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentutilizationasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentutilizationasyncaggregatequeryresponse

    if AgentutilizationasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentutilizationasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Agentutilizationaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Agentutilizationaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

