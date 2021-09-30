package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SecuresessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SecuresessionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Securesession
type Securesession struct { 
    


    // Flow - The flow to execute securely
    Flow Domainentityref `json:"flow"`


    // UserData - Customer-provided data
    UserData string `json:"userData"`


    // State - The current state of a secure session
    State string `json:"state"`


    // SourceParticipantId - Unique identifier for the participant initiating the secure session.
    SourceParticipantId string `json:"sourceParticipantId"`


    // Disconnect - If true, disconnect the agent after creating the session
    Disconnect bool `json:"disconnect"`


    

}

// String returns a JSON representation of the model
func (o *Securesession) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Securesession) MarshalJSON() ([]byte, error) {
    type Alias Securesession

    if SecuresessionMarshalled {
        return []byte("{}"), nil
    }
    SecuresessionMarshalled = true

    return json.Marshal(&struct { 
        
        
        Flow Domainentityref `json:"flow"`
        
        UserData string `json:"userData"`
        
        State string `json:"state"`
        
        SourceParticipantId string `json:"sourceParticipantId"`
        
        Disconnect bool `json:"disconnect"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

