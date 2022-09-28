package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationconfigurationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Integrationconfiguration - Configuration for an Integration
type Integrationconfiguration struct { 
    


    // Name - The name of the integration, used to distinguish this integration from others of the same type.
    Name string `json:"name"`


    // Version - Version number required for updates.
    Version int `json:"version"`


    // Properties - Key-value configuration settings described by the schema in the propertiesSchemaUri field.
    Properties interface{} `json:"properties"`


    // Advanced - Advanced configuration described by the schema in the advancedSchemaUri field.
    Advanced interface{} `json:"advanced"`


    // Notes - Notes about the integration.
    Notes string `json:"notes"`


    // Credentials - Credentials required by the integration. The required keys are indicated in the credentials property of the Integration Type
    Credentials map[string]Integrationconfigurationcredential `json:"credentials"`


    

}

// String returns a JSON representation of the model
func (o *Integrationconfiguration) String() string {
    
    
     o.Properties = Interface{} 
     o.Advanced = Interface{} 
    
     o.Credentials = map[string]Integrationconfigurationcredential{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Integrationconfiguration

    if IntegrationconfigurationMarshalled {
        return []byte("{}"), nil
    }
    IntegrationconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Properties interface{} `json:"properties"`
        
        Advanced interface{} `json:"advanced"`
        
        Notes string `json:"notes"`
        
        Credentials map[string]Integrationconfigurationcredential `json:"credentials"`
        *Alias
    }{

        


        


        


        
        Properties: Interface{},
        


        
        Advanced: Interface{},
        


        


        
        Credentials: map[string]Integrationconfigurationcredential{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

