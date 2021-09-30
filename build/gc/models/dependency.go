package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DependencyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DependencyDud struct { 
    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dependency
type Dependency struct { 
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


    

}

// String returns a JSON representation of the model
func (o *Dependency) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dependency) MarshalJSON() ([]byte, error) {
    type Alias Dependency

    if DependencyMarshalled {
        return []byte("{}"), nil
    }
    DependencyMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Version string `json:"version"`
        
        VarType string `json:"type"`
        
        Deleted bool `json:"deleted"`
        
        Updated bool `json:"updated"`
        
        StateUnknown bool `json:"stateUnknown"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

