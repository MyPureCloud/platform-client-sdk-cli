package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentactionpayloadMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentactionpayloadDud struct { 
    


    

}

// Agentactionpayload
type Agentactionpayload struct { 
    // AfterCallWork - Boolean flag to indicate if the agent action is in ACW stage.
    AfterCallWork bool `json:"afterCallWork"`


    // ChecklistItems - The list of checklist items and the agent action taken on each one of them.
    ChecklistItems []Agentaction `json:"checklistItems"`

}

// String returns a JSON representation of the model
func (o *Agentactionpayload) String() string {
    
     o.ChecklistItems = []Agentaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentactionpayload) MarshalJSON() ([]byte, error) {
    type Alias Agentactionpayload

    if AgentactionpayloadMarshalled {
        return []byte("{}"), nil
    }
    AgentactionpayloadMarshalled = true

    return json.Marshal(&struct {
        
        AfterCallWork bool `json:"afterCallWork"`
        
        ChecklistItems []Agentaction `json:"checklistItems"`
        *Alias
    }{

        


        
        ChecklistItems: []Agentaction{{}},
        

        Alias: (*Alias)(u),
    })
}

