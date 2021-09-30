package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DependencyobjectMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DependencyobjectDud struct { 
    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dependencyobject
type Dependencyobject struct { 
    // Id - The dependency identifier
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Version
    Version string `json:"version"`


    // VarType
    VarType string `json:"type"`


    // Deleted
    Deleted bool `json:"deleted"`


    // Updated
    Updated bool `json:"updated"`


    // StateUnknown
    StateUnknown bool `json:"stateUnknown"`


    // ConsumedResources
    ConsumedResources []Dependency `json:"consumedResources"`


    // ConsumingResources
    ConsumingResources []Dependency `json:"consumingResources"`


    

}

// String returns a JSON representation of the model
func (o *Dependencyobject) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ConsumedResources = []Dependency{{}} 
    
    
    
     o.ConsumingResources = []Dependency{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dependencyobject) MarshalJSON() ([]byte, error) {
    type Alias Dependencyobject

    if DependencyobjectMarshalled {
        return []byte("{}"), nil
    }
    DependencyobjectMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Version string `json:"version"`
        
        VarType string `json:"type"`
        
        Deleted bool `json:"deleted"`
        
        Updated bool `json:"updated"`
        
        StateUnknown bool `json:"stateUnknown"`
        
        ConsumedResources []Dependency `json:"consumedResources"`
        
        ConsumingResources []Dependency `json:"consumingResources"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        ConsumedResources: []Dependency{{}},
        

        

        
        ConsumingResources: []Dependency{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

