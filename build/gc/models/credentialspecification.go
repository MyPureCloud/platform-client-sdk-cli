package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CredentialspecificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CredentialspecificationDud struct { 
    Required bool `json:"required"`


    Title string `json:"title"`


    CredentialTypes []string `json:"credentialTypes"`

}

// Credentialspecification - Specifies the requirements for a credential that can be provided for configuring an integration
type Credentialspecification struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Credentialspecification) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Credentialspecification) MarshalJSON() ([]byte, error) {
    type Alias Credentialspecification

    if CredentialspecificationMarshalled {
        return []byte("{}"), nil
    }
    CredentialspecificationMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

