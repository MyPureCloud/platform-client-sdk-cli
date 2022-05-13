package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentDud struct { 
    

}

// Agent
type Agent struct { 
    // Stage - The current stage for this agent
    Stage string `json:"stage"`

}

// String returns a JSON representation of the model
func (o *Agent) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agent) MarshalJSON() ([]byte, error) {
    type Alias Agent

    if AgentMarshalled {
        return []byte("{}"), nil
    }
    AgentMarshalled = true

    return json.Marshal(&struct {
        
        Stage string `json:"stage"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

