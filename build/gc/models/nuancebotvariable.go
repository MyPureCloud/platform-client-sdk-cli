package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancebotvariableMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancebotvariableDud struct { 
    


    


    


    


    


    

}

// Nuancebotvariable - Model for a Nuance bot variable
type Nuancebotvariable struct { 
    // Id - The variable ID
    Id string `json:"id"`


    // Name - The variable display name
    Name string `json:"name"`


    // Description - The variable description
    Description string `json:"description"`


    // Reserved - True if the variable is a reserved variable
    Reserved bool `json:"reserved"`


    // SimpleVariableInfo - The type information for this variable
    SimpleVariableInfo string `json:"simpleVariableInfo"`


    // ComplexGenericVariableInfo - The type information for this variable
    ComplexGenericVariableInfo Complexvariableinfo `json:"complexGenericVariableInfo"`

}

// String returns a JSON representation of the model
func (o *Nuancebotvariable) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancebotvariable) MarshalJSON() ([]byte, error) {
    type Alias Nuancebotvariable

    if NuancebotvariableMarshalled {
        return []byte("{}"), nil
    }
    NuancebotvariableMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Reserved bool `json:"reserved"`
        
        SimpleVariableInfo string `json:"simpleVariableInfo"`
        
        ComplexGenericVariableInfo Complexvariableinfo `json:"complexGenericVariableInfo"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

