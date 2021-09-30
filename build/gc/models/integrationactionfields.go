package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationactionfieldsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationactionfieldsDud struct { 
    


    

}

// Integrationactionfields
type Integrationactionfields struct { 
    // IntegrationAction - Reference to the Integration Action to be used when integrationAction type is qualified
    IntegrationAction Integrationaction `json:"integrationAction"`


    // RequestMappings - Collection of Request Mappings to use
    RequestMappings []Requestmapping `json:"requestMappings"`

}

// String returns a JSON representation of the model
func (o *Integrationactionfields) String() string {
    
    
    
    
    
    
     o.RequestMappings = []Requestmapping{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationactionfields) MarshalJSON() ([]byte, error) {
    type Alias Integrationactionfields

    if IntegrationactionfieldsMarshalled {
        return []byte("{}"), nil
    }
    IntegrationactionfieldsMarshalled = true

    return json.Marshal(&struct { 
        IntegrationAction Integrationaction `json:"integrationAction"`
        
        RequestMappings []Requestmapping `json:"requestMappings"`
        
        *Alias
    }{
        

        

        

        
        RequestMappings: []Requestmapping{{}},
        

        
        Alias: (*Alias)(u),
    })
}

