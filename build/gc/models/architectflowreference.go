package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArchitectflowreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArchitectflowreferenceDud struct { 
    


    


    


    

}

// Architectflowreference
type Architectflowreference struct { 
    // Id - The flow identifier.
    Id string `json:"id"`


    // Name - The flow name.
    Name string `json:"name"`


    // VarType - The flow type.
    VarType string `json:"type"`


    // Version - The flow version.
    Version string `json:"version"`

}

// String returns a JSON representation of the model
func (o *Architectflowreference) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Architectflowreference) MarshalJSON() ([]byte, error) {
    type Alias Architectflowreference

    if ArchitectflowreferenceMarshalled {
        return []byte("{}"), nil
    }
    ArchitectflowreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Version string `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

