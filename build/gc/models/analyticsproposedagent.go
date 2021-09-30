package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsproposedagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsproposedagentDud struct { 
    


    

}

// Analyticsproposedagent
type Analyticsproposedagent struct { 
    // AgentRank - Proposed agent rank for this conversation from predictive routing (lower is better)
    AgentRank int `json:"agentRank"`


    // ProposedAgentId - Unique identifier for the agent that was proposed by predictive routing
    ProposedAgentId string `json:"proposedAgentId"`

}

// String returns a JSON representation of the model
func (o *Analyticsproposedagent) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsproposedagent) MarshalJSON() ([]byte, error) {
    type Alias Analyticsproposedagent

    if AnalyticsproposedagentMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsproposedagentMarshalled = true

    return json.Marshal(&struct { 
        AgentRank int `json:"agentRank"`
        
        ProposedAgentId string `json:"proposedAgentId"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

