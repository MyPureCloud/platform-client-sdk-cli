package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateverifierresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateverifierresponseDud struct { 
    


    


    


    


    


    

}

// Createverifierresponse
type Createverifierresponse struct { 
    // Id - The unique identifier of the verifier.
    Id string `json:"id"`


    // Name - The name of the verifier.
    Name string `json:"name"`


    // VarType - The type of verifier.
    VarType string `json:"type"`


    // Enabled - Indicates whether this verifier is enabled.
    Enabled bool `json:"enabled"`


    // KeyUri - The key URI for TOTP authenticator app registration.
    KeyUri string `json:"keyUri"`


    // VarDefault - Indicates whether this is the default verifier.
    VarDefault bool `json:"default"`

}

// String returns a JSON representation of the model
func (o *Createverifierresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createverifierresponse) MarshalJSON() ([]byte, error) {
    type Alias Createverifierresponse

    if CreateverifierresponseMarshalled {
        return []byte("{}"), nil
    }
    CreateverifierresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Enabled bool `json:"enabled"`
        
        KeyUri string `json:"keyUri"`
        
        VarDefault bool `json:"default"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

