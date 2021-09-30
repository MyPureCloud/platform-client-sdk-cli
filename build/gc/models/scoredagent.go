package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScoredagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScoredagentDud struct { 
    


    

}

// Scoredagent
type Scoredagent struct { 
    // Agent - The agent
    Agent Addressableentityref `json:"agent"`


    // Score - Agent's score for the current conversation, from 0 - 100, higher being better
    Score int `json:"score"`

}

// String returns a JSON representation of the model
func (o *Scoredagent) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scoredagent) MarshalJSON() ([]byte, error) {
    type Alias Scoredagent

    if ScoredagentMarshalled {
        return []byte("{}"), nil
    }
    ScoredagentMarshalled = true

    return json.Marshal(&struct { 
        Agent Addressableentityref `json:"agent"`
        
        Score int `json:"score"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

