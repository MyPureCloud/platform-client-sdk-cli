package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentsintegrationslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentsintegrationslistingDud struct { 
    

}

// Agentsintegrationslisting
type Agentsintegrationslisting struct { 
    // Entities
    Entities []Agentintegrationsresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Agentsintegrationslisting) String() string {
     o.Entities = []Agentintegrationsresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentsintegrationslisting) MarshalJSON() ([]byte, error) {
    type Alias Agentsintegrationslisting

    if AgentsintegrationslistingMarshalled {
        return []byte("{}"), nil
    }
    AgentsintegrationslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Agentintegrationsresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Agentintegrationsresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

