package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchintegrationactionfieldsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchintegrationactionfieldsDud struct { 
    


    

}

// Patchintegrationactionfields
type Patchintegrationactionfields struct { 
    // IntegrationAction - Reference to the Integration Action to be used when integrationAction type is qualified
    IntegrationAction Patchintegrationaction `json:"integrationAction"`


    // RequestMappings - Collection of Request Mappings to use
    RequestMappings []Requestmapping `json:"requestMappings"`

}

// String returns a JSON representation of the model
func (o *Patchintegrationactionfields) String() string {
    
    
    
    
    
    
     o.RequestMappings = []Requestmapping{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchintegrationactionfields) MarshalJSON() ([]byte, error) {
    type Alias Patchintegrationactionfields

    if PatchintegrationactionfieldsMarshalled {
        return []byte("{}"), nil
    }
    PatchintegrationactionfieldsMarshalled = true

    return json.Marshal(&struct { 
        IntegrationAction Patchintegrationaction `json:"integrationAction"`
        
        RequestMappings []Requestmapping `json:"requestMappings"`
        
        *Alias
    }{
        

        

        

        
        RequestMappings: []Requestmapping{{}},
        

        
        Alias: (*Alias)(u),
    })
}

