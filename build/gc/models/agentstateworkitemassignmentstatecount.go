package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateworkitemassignmentstatecountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateworkitemassignmentstatecountDud struct { 
    


    

}

// Agentstateworkitemassignmentstatecount
type Agentstateworkitemassignmentstatecount struct { 
    // AssignmentState - Workitem assignment state
    AssignmentState string `json:"assignmentState"`


    // Count - Count of workitems with this assignment state
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Agentstateworkitemassignmentstatecount) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateworkitemassignmentstatecount) MarshalJSON() ([]byte, error) {
    type Alias Agentstateworkitemassignmentstatecount

    if AgentstateworkitemassignmentstatecountMarshalled {
        return []byte("{}"), nil
    }
    AgentstateworkitemassignmentstatecountMarshalled = true

    return json.Marshal(&struct {
        
        AssignmentState string `json:"assignmentState"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

