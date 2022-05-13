package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentmanagementunitreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentmanagementunitreferenceDud struct { 
    


    


    

}

// Agentmanagementunitreference - A reference from agent to management unit
type Agentmanagementunitreference struct { 
    // User - The user (agent) for whom the management unit was requested
    User Userreference `json:"user"`


    // ManagementUnit - The management to which the user (agent) belongs
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // BusinessUnit - The business unit to which the user (agent) belongs. Populate with expand=businessUnit
    BusinessUnit Businessunitreference `json:"businessUnit"`

}

// String returns a JSON representation of the model
func (o *Agentmanagementunitreference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentmanagementunitreference) MarshalJSON() ([]byte, error) {
    type Alias Agentmanagementunitreference

    if AgentmanagementunitreferenceMarshalled {
        return []byte("{}"), nil
    }
    AgentmanagementunitreferenceMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

