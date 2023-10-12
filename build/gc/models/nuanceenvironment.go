package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuanceenvironmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuanceenvironmentDud struct { 
    


    


    


    

}

// Nuanceenvironment - Model for a Nuance bot environment
type Nuanceenvironment struct { 
    // Id - The environment ID
    Id string `json:"id"`


    // Name - The environment name
    Name string `json:"name"`


    // EnvironmentDesignation - The environment type, usually a value like SANDBOX or PRODUCTION
    EnvironmentDesignation string `json:"environmentDesignation"`


    // ExecutionHost - The host used to execute this bot, similar to 'api.nuance.com:443'
    ExecutionHost string `json:"executionHost"`

}

// String returns a JSON representation of the model
func (o *Nuanceenvironment) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuanceenvironment) MarshalJSON() ([]byte, error) {
    type Alias Nuanceenvironment

    if NuanceenvironmentMarshalled {
        return []byte("{}"), nil
    }
    NuanceenvironmentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        EnvironmentDesignation string `json:"environmentDesignation"`
        
        ExecutionHost string `json:"executionHost"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

