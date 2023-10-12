package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancebottransfernodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancebottransfernodeDud struct { 
    


    


    


    


    

}

// Nuancebottransfernode - Model for a Nuance bot transfer node
type Nuancebottransfernode struct { 
    // Id - The transfer node ID
    Id string `json:"id"`


    // Name - The transfer node name
    Name string `json:"name"`


    // VarType - The transfer node type
    VarType string `json:"type"`


    // Description - The transfer node description
    Description string `json:"description"`


    // RequestVariables - List of variables associated with this transfer node
    RequestVariables []Nuancebotvariable `json:"requestVariables"`

}

// String returns a JSON representation of the model
func (o *Nuancebottransfernode) String() string {
    
    
    
    
     o.RequestVariables = []Nuancebotvariable{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancebottransfernode) MarshalJSON() ([]byte, error) {
    type Alias Nuancebottransfernode

    if NuancebottransfernodeMarshalled {
        return []byte("{}"), nil
    }
    NuancebottransfernodeMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Description string `json:"description"`
        
        RequestVariables []Nuancebotvariable `json:"requestVariables"`
        *Alias
    }{

        


        


        


        


        
        RequestVariables: []Nuancebotvariable{{}},
        

        Alias: (*Alias)(u),
    })
}

