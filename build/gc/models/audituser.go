package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AudituserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AudituserDud struct { 
    


    


    

}

// Audituser
type Audituser struct { 
    // Id - The ID (UUID) of the user who initiated the action of this AuditMessage.
    Id string `json:"id"`


    // Name - The full username of the user who initiated the action of this AuditMessage.
    Name string `json:"name"`


    // Display - The display name of the user who initiated the action of this AuditMessage.
    Display string `json:"display"`

}

// String returns a JSON representation of the model
func (o *Audituser) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Audituser) MarshalJSON() ([]byte, error) {
    type Alias Audituser

    if AudituserMarshalled {
        return []byte("{}"), nil
    }
    AudituserMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Display string `json:"display"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

