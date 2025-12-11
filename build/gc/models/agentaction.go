package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentactionDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentaction
type Agentaction struct { 
    // Id - ID of the checklist item.
    Id string `json:"id"`


    // Name - Name of the checklist item.
    Name string `json:"name"`


    // AgentAction - The agent action taken on a checklist item.
    AgentAction string `json:"agentAction"`


    

}

// String returns a JSON representation of the model
func (o *Agentaction) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentaction) MarshalJSON() ([]byte, error) {
    type Alias Agentaction

    if AgentactionMarshalled {
        return []byte("{}"), nil
    }
    AgentactionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        AgentAction string `json:"agentAction"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

