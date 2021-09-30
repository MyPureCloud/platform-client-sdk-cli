package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimgenesysuserexternalidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimgenesysuserexternalidDud struct { 
    


    

}

// Scimgenesysuserexternalid - External Identifiers of user. The external identifier must be unique within the organization and the 'authority'
type Scimgenesysuserexternalid struct { 
    // Authority - Authority, or scope, of \"externalId\". Allows multiple external identifiers to be defined. Represents the source of the external identifier.
    Authority string `json:"authority"`


    // Value - Identifier of the user in an external system.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Scimgenesysuserexternalid) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimgenesysuserexternalid) MarshalJSON() ([]byte, error) {
    type Alias Scimgenesysuserexternalid

    if ScimgenesysuserexternalidMarshalled {
        return []byte("{}"), nil
    }
    ScimgenesysuserexternalidMarshalled = true

    return json.Marshal(&struct { 
        Authority string `json:"authority"`
        
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

