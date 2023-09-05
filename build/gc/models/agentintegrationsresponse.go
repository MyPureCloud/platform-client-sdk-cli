package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentintegrationsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentintegrationsresponseDud struct { 
    


    


    


    

}

// Agentintegrationsresponse
type Agentintegrationsresponse struct { 
    // Agent - The user associated with the integrations
    Agent Userreference `json:"agent"`


    // SelectedIntegration - The integration selected for the agent. If not set, no integration will be used for the agent
    SelectedIntegration Wfmintegrationreference `json:"selectedIntegration"`


    // UserSelected - Whether the integration association has been manually selected
    UserSelected bool `json:"userSelected"`


    // AssociatedIntegrations - The list of integrations associated with the agent
    AssociatedIntegrations []Agentintegrationassociationresponse `json:"associatedIntegrations"`

}

// String returns a JSON representation of the model
func (o *Agentintegrationsresponse) String() string {
    
    
    
     o.AssociatedIntegrations = []Agentintegrationassociationresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentintegrationsresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentintegrationsresponse

    if AgentintegrationsresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentintegrationsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Agent Userreference `json:"agent"`
        
        SelectedIntegration Wfmintegrationreference `json:"selectedIntegration"`
        
        UserSelected bool `json:"userSelected"`
        
        AssociatedIntegrations []Agentintegrationassociationresponse `json:"associatedIntegrations"`
        *Alias
    }{

        


        


        


        
        AssociatedIntegrations: []Agentintegrationassociationresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

