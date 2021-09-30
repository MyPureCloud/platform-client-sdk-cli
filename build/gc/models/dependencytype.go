package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DependencytypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DependencytypeDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Dependencytype
type Dependencytype struct { 
    // Id - The dependency type identifier
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Versioned
    Versioned bool `json:"versioned"`


    

}

// String returns a JSON representation of the model
func (o *Dependencytype) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dependencytype) MarshalJSON() ([]byte, error) {
    type Alias Dependencytype

    if DependencytypeMarshalled {
        return []byte("{}"), nil
    }
    DependencytypeMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Versioned bool `json:"versioned"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

