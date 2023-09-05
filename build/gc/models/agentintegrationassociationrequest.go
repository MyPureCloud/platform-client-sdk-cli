package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentintegrationassociationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentintegrationassociationrequestDud struct { 
    


    


    

}

// Agentintegrationassociationrequest
type Agentintegrationassociationrequest struct { 
    // AgentExternalId - The ID of the agent in external system
    AgentExternalId string `json:"agentExternalId"`


    // IntegrationId - The ID of the integration associated with the agent
    IntegrationId string `json:"integrationId"`


    // Locked - Whether agentExternalId should be protected from update by automatic processes
    Locked bool `json:"locked"`

}

// String returns a JSON representation of the model
func (o *Agentintegrationassociationrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentintegrationassociationrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentintegrationassociationrequest

    if AgentintegrationassociationrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentintegrationassociationrequestMarshalled = true

    return json.Marshal(&struct {
        
        AgentExternalId string `json:"agentExternalId"`
        
        IntegrationId string `json:"integrationId"`
        
        Locked bool `json:"locked"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

