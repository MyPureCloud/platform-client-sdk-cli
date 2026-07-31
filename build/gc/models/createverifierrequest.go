package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateverifierrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateverifierrequestDud struct { 
    


    


    


    


    


    


    

}

// Createverifierrequest
type Createverifierrequest struct { 
    // Algorithm - The hashing algorithm for the TOTP verifier.
    Algorithm string `json:"algorithm"`


    // Digits - The number of digits in the TOTP code. Must be between 6 and 12.
    Digits int `json:"digits"`


    // Enabled - Indicates whether this verifier will be enabled.
    Enabled bool `json:"enabled"`


    // Name - The name of the verifier. Maximum length is 100 characters.
    Name string `json:"name"`


    // Period - The time period in seconds for the TOTP code.
    Period int `json:"period"`


    // SecretSize - The size of the shared secret in bytes. Must be between 10 and 64.
    SecretSize int `json:"secretSize"`


    // VarDefault - Indicates whether this will be the default verifier.
    VarDefault bool `json:"default"`

}

// String returns a JSON representation of the model
func (o *Createverifierrequest) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createverifierrequest) MarshalJSON() ([]byte, error) {
    type Alias Createverifierrequest

    if CreateverifierrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateverifierrequestMarshalled = true

    return json.Marshal(&struct {
        
        Algorithm string `json:"algorithm"`
        
        Digits int `json:"digits"`
        
        Enabled bool `json:"enabled"`
        
        Name string `json:"name"`
        
        Period int `json:"period"`
        
        SecretSize int `json:"secretSize"`
        
        VarDefault bool `json:"default"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

