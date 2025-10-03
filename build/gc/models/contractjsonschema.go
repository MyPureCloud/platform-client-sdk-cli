package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContractjsonschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContractjsonschemaDud struct { 
    


    


    


    


    


    


    


    


    

}

// Contractjsonschema - A JSON Schema document.
type Contractjsonschema struct { 
    // Id
    Id string `json:"id"`


    // Schema
    Schema string `json:"$schema"`


    // Title
    Title string `json:"title"`


    // Description
    Description string `json:"description"`


    // VarType
    VarType string `json:"type"`


    // Required
    Required []string `json:"required"`


    // Properties
    Properties map[string]interface{} `json:"properties"`


    // AdditionalProperties
    AdditionalProperties interface{} `json:"additionalProperties"`


    // Definitions
    Definitions map[string]Contractdefinition `json:"definitions"`

}

// String returns a JSON representation of the model
func (o *Contractjsonschema) String() string {
    
    
    
    
    
     o.Required = []string{""} 
     o.Properties = map[string]interface{}{"": Interface{}} 
     o.AdditionalProperties = Interface{} 
     o.Definitions = map[string]Contractdefinition{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contractjsonschema) MarshalJSON() ([]byte, error) {
    type Alias Contractjsonschema

    if ContractjsonschemaMarshalled {
        return []byte("{}"), nil
    }
    ContractjsonschemaMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Schema string `json:"$schema"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        
        Required []string `json:"required"`
        
        Properties map[string]interface{} `json:"properties"`
        
        AdditionalProperties interface{} `json:"additionalProperties"`
        
        Definitions map[string]Contractdefinition `json:"definitions"`
        *Alias
    }{

        


        


        


        


        


        
        Required: []string{""},
        


        
        Properties: map[string]interface{}{"": Interface{}},
        


        
        AdditionalProperties: Interface{},
        


        
        Definitions: map[string]Contractdefinition{"": {}},
        

        Alias: (*Alias)(u),
    })
}

