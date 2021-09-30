package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimserviceproviderconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimserviceproviderconfigDud struct { 
    Schemas []string `json:"schemas"`


    DocumentationUri string `json:"documentationUri"`


    Patch Scimserviceproviderconfigsimplefeature `json:"patch"`


    Filter Scimserviceproviderconfigfilterfeature `json:"filter"`


    Etag Scimserviceproviderconfigsimplefeature `json:"etag"`


    Sort Scimserviceproviderconfigsimplefeature `json:"sort"`


    Bulk Scimserviceproviderconfigbulkfeature `json:"bulk"`


    ChangePassword Scimserviceproviderconfigsimplefeature `json:"changePassword"`


    AuthenticationSchemes []Scimserviceproviderconfigauthenticationscheme `json:"authenticationSchemes"`


    Meta Scimmetadata `json:"meta"`

}

// Scimserviceproviderconfig - Defines a SCIM service provider's configuration.
type Scimserviceproviderconfig struct { 
    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Scimserviceproviderconfig) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimserviceproviderconfig) MarshalJSON() ([]byte, error) {
    type Alias Scimserviceproviderconfig

    if ScimserviceproviderconfigMarshalled {
        return []byte("{}"), nil
    }
    ScimserviceproviderconfigMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

