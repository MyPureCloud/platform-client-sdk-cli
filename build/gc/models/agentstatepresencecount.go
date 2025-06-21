package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstatepresencecountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstatepresencecountDud struct { 
    


    


    

}

// Agentstatepresencecount
type Agentstatepresencecount struct { 
    // SystemPresence - System presence
    SystemPresence string `json:"systemPresence"`


    // OrganizationPresenceId - The identifier for the organization presence
    OrganizationPresenceId string `json:"organizationPresenceId"`


    // Count - Count of users with this system presence and organization presence
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Agentstatepresencecount) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstatepresencecount) MarshalJSON() ([]byte, error) {
    type Alias Agentstatepresencecount

    if AgentstatepresencecountMarshalled {
        return []byte("{}"), nil
    }
    AgentstatepresencecountMarshalled = true

    return json.Marshal(&struct {
        
        SystemPresence string `json:"systemPresence"`
        
        OrganizationPresenceId string `json:"organizationPresenceId"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

