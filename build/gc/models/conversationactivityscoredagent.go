package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationactivityscoredagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationactivityscoredagentDud struct { 
    


    

}

// Conversationactivityscoredagent
type Conversationactivityscoredagent struct { 
    // AgentScore - Assigned agent score for this conversation (0 - 100, higher being better)
    AgentScore int `json:"agentScore"`


    // ScoredAgentId - Unique identifier for the agent that was scored for this conversation
    ScoredAgentId string `json:"scoredAgentId"`

}

// String returns a JSON representation of the model
func (o *Conversationactivityscoredagent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationactivityscoredagent) MarshalJSON() ([]byte, error) {
    type Alias Conversationactivityscoredagent

    if ConversationactivityscoredagentMarshalled {
        return []byte("{}"), nil
    }
    ConversationactivityscoredagentMarshalled = true

    return json.Marshal(&struct {
        
        AgentScore int `json:"agentScore"`
        
        ScoredAgentId string `json:"scoredAgentId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

