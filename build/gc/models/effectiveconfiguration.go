package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EffectiveconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EffectiveconfigurationDud struct { 
    


    


    


    


    

}

// Effectiveconfiguration - Effective Configuration for an ClientApp. This is comprised of the integration specific configuration along with overrides specified in the integration type.
type Effectiveconfiguration struct { 
    // Properties - Key-value configuration settings described by the schema in the propertiesSchemaUri field.
    Properties map[string]interface{} `json:"properties"`


    // Advanced - Advanced configuration described by the schema in the advancedSchemaUri field.
    Advanced map[string]interface{} `json:"advanced"`


    // Name - The name of the integration, used to distinguish this integration from others of the same type.
    Name string `json:"name"`


    // Notes - Notes about the integration.
    Notes string `json:"notes"`


    // Credentials - Credentials required by the integration. The required keys are indicated in the credentials property of the Integration Type
    Credentials map[string]Credentialinfo `json:"credentials"`

}

// String returns a JSON representation of the model
func (o *Effectiveconfiguration) String() string {
     o.Properties = map[string]interface{}{"": Interface{}} 
     o.Advanced = map[string]interface{}{"": Interface{}} 
    
    
     o.Credentials = map[string]Credentialinfo{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Effectiveconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Effectiveconfiguration

    if EffectiveconfigurationMarshalled {
        return []byte("{}"), nil
    }
    EffectiveconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Properties map[string]interface{} `json:"properties"`
        
        Advanced map[string]interface{} `json:"advanced"`
        
        Name string `json:"name"`
        
        Notes string `json:"notes"`
        
        Credentials map[string]Credentialinfo `json:"credentials"`
        *Alias
    }{

        
        Properties: map[string]interface{}{"": Interface{}},
        


        
        Advanced: map[string]interface{}{"": Interface{}},
        


        


        


        
        Credentials: map[string]Credentialinfo{"": {}},
        

        Alias: (*Alias)(u),
    })
}

