package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CredentialdescriptorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CredentialdescriptorDud struct { 
    


    


    

}

// Credentialdescriptor
type Credentialdescriptor struct { 
    // VarType - The credential type (e.g., 'public-key').
    VarType string `json:"type"`


    // Id - The credential identifier (base64url-encoded).
    Id string `json:"id"`


    // Transports - Hints regarding which transports the credential supports.
    Transports []string `json:"transports"`

}

// String returns a JSON representation of the model
func (o *Credentialdescriptor) String() string {
    
    
     o.Transports = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Credentialdescriptor) MarshalJSON() ([]byte, error) {
    type Alias Credentialdescriptor

    if CredentialdescriptorMarshalled {
        return []byte("{}"), nil
    }
    CredentialdescriptorMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Id string `json:"id"`
        
        Transports []string `json:"transports"`
        *Alias
    }{

        


        


        
        Transports: []string{""},
        

        Alias: (*Alias)(u),
    })
}

