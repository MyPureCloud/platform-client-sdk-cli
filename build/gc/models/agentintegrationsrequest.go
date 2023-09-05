package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentintegrationsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentintegrationsrequestDud struct { 
    


    

}

// Agentintegrationsrequest
type Agentintegrationsrequest struct { 
    // SelectedIntegrationId - The ID of the integration selected for the agent. If not set, no integration will be used for the agent
    SelectedIntegrationId string `json:"selectedIntegrationId"`


    // AssociatedIntegrations - The list of integrations associated with the agent
    AssociatedIntegrations []Agentintegrationassociationrequest `json:"associatedIntegrations"`

}

// String returns a JSON representation of the model
func (o *Agentintegrationsrequest) String() string {
    
     o.AssociatedIntegrations = []Agentintegrationassociationrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentintegrationsrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentintegrationsrequest

    if AgentintegrationsrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentintegrationsrequestMarshalled = true

    return json.Marshal(&struct {
        
        SelectedIntegrationId string `json:"selectedIntegrationId"`
        
        AssociatedIntegrations []Agentintegrationassociationrequest `json:"associatedIntegrations"`
        *Alias
    }{

        


        
        AssociatedIntegrations: []Agentintegrationassociationrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

