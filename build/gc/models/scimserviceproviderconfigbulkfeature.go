package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimserviceproviderconfigbulkfeatureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimserviceproviderconfigbulkfeatureDud struct { 
    Supported bool `json:"supported"`


    MaxOperations int `json:"maxOperations"`


    MaxPayloadSize int `json:"maxPayloadSize"`

}

// Scimserviceproviderconfigbulkfeature - Defines a \"bulk\" request in the SCIM service provider's configuration.
type Scimserviceproviderconfigbulkfeature struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfigbulkfeature) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimserviceproviderconfigbulkfeature) MarshalJSON() ([]byte, error) {
    type Alias Scimserviceproviderconfigbulkfeature

    if ScimserviceproviderconfigbulkfeatureMarshalled {
        return []byte("{}"), nil
    }
    ScimserviceproviderconfigbulkfeatureMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

