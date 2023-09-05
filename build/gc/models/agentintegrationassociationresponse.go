package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentintegrationassociationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentintegrationassociationresponseDud struct { 
    


    


    

}

// Agentintegrationassociationresponse
type Agentintegrationassociationresponse struct { 
    // AgentExternalId - ID of the agent in external system
    AgentExternalId string `json:"agentExternalId"`


    // Integration - The integration associated with the agent
    Integration Wfmintegrationreference `json:"integration"`


    // Locked - Whether agentExternalId should be protected from update by automatic processes
    Locked bool `json:"locked"`

}

// String returns a JSON representation of the model
func (o *Agentintegrationassociationresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentintegrationassociationresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentintegrationassociationresponse

    if AgentintegrationassociationresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentintegrationassociationresponseMarshalled = true

    return json.Marshal(&struct {
        
        AgentExternalId string `json:"agentExternalId"`
        
        Integration Wfmintegrationreference `json:"integration"`
        
        Locked bool `json:"locked"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

