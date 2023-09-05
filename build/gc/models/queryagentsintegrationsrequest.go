package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryagentsintegrationsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryagentsintegrationsrequestDud struct { 
    

}

// Queryagentsintegrationsrequest - Query integrations for agents request
type Queryagentsintegrationsrequest struct { 
    // AgentIds - The IDs of the agents to query associated integrations
    AgentIds []string `json:"agentIds"`

}

// String returns a JSON representation of the model
func (o *Queryagentsintegrationsrequest) String() string {
     o.AgentIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryagentsintegrationsrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryagentsintegrationsrequest

    if QueryagentsintegrationsrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryagentsintegrationsrequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentIds []string `json:"agentIds"`
        *Alias
    }{

        
        AgentIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

