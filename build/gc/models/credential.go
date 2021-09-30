package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CredentialMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CredentialDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Credential
type Credential struct { 
    


    // Name
    Name string `json:"name"`


    // VarType - The type of credential.
    VarType Credentialtype `json:"type"`


    // CredentialFields
    CredentialFields map[string]string `json:"credentialFields"`


    

}

// String returns a JSON representation of the model
func (o *Credential) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.CredentialFields = map[string]string{"": ""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Credential) MarshalJSON() ([]byte, error) {
    type Alias Credential

    if CredentialMarshalled {
        return []byte("{}"), nil
    }
    CredentialMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        VarType Credentialtype `json:"type"`
        
        CredentialFields map[string]string `json:"credentialFields"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        CredentialFields: map[string]string{"": ""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

