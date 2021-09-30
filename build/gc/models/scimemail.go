package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimemailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimemailDud struct { 
    


    


    

}

// Scimemail - Defines a SCIM email address.
type Scimemail struct { 
    // Value - The email address. Is immutable if \"type\" is set to \"other\".
    Value string `json:"value"`


    // VarType - The type of email address. \"value\" is immutable if \"type\" is set to \"other\".
    VarType string `json:"type"`


    // Primary - Indicates whether the email address is the primary email address.
    Primary bool `json:"primary"`

}

// String returns a JSON representation of the model
func (o *Scimemail) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimemail) MarshalJSON() ([]byte, error) {
    type Alias Scimemail

    if ScimemailMarshalled {
        return []byte("{}"), nil
    }
    ScimemailMarshalled = true

    return json.Marshal(&struct { 
        Value string `json:"value"`
        
        VarType string `json:"type"`
        
        Primary bool `json:"primary"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

