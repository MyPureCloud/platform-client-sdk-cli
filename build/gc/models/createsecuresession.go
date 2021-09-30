package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatesecuresessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatesecuresessionDud struct { 
    


    


    


    

}

// Createsecuresession
type Createsecuresession struct { 
    // SourceParticipantId - requesting participant
    SourceParticipantId string `json:"sourceParticipantId"`


    // FlowId - the flow id to execute in the secure session
    FlowId string `json:"flowId"`


    // UserData - user data for the secure session
    UserData string `json:"userData"`


    // Disconnect - if true, disconnect the agent after creating the session
    Disconnect bool `json:"disconnect"`

}

// String returns a JSON representation of the model
func (o *Createsecuresession) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createsecuresession) MarshalJSON() ([]byte, error) {
    type Alias Createsecuresession

    if CreatesecuresessionMarshalled {
        return []byte("{}"), nil
    }
    CreatesecuresessionMarshalled = true

    return json.Marshal(&struct { 
        SourceParticipantId string `json:"sourceParticipantId"`
        
        FlowId string `json:"flowId"`
        
        UserData string `json:"userData"`
        
        Disconnect bool `json:"disconnect"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

