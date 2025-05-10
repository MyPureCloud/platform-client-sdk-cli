package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JsonschemawithdefinitionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JsonschemawithdefinitionsDud struct { 
    


    


    


    


    


    


    


    


    

}

// Jsonschemawithdefinitions - A JSON Schema document.
type Jsonschemawithdefinitions struct { 
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
    Definitions map[string]Definition `json:"definitions"`

}

// String returns a JSON representation of the model
func (o *Jsonschemawithdefinitions) String() string {
    
    
    
    
    
     o.Required = []string{""} 
     o.Properties = map[string]interface{}{"": Interface{}} 
     o.AdditionalProperties = Interface{} 
     o.Definitions = map[string]Definition{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Jsonschemawithdefinitions) MarshalJSON() ([]byte, error) {
    type Alias Jsonschemawithdefinitions

    if JsonschemawithdefinitionsMarshalled {
        return []byte("{}"), nil
    }
    JsonschemawithdefinitionsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Schema string `json:"$schema"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        
        Required []string `json:"required"`
        
        Properties map[string]interface{} `json:"properties"`
        
        AdditionalProperties interface{} `json:"additionalProperties"`
        
        Definitions map[string]Definition `json:"definitions"`
        *Alias
    }{

        


        


        


        


        


        
        Required: []string{""},
        


        
        Properties: map[string]interface{}{"": Interface{}},
        


        
        AdditionalProperties: Interface{},
        


        
        Definitions: map[string]Definition{"": {}},
        

        Alias: (*Alias)(u),
    })
}

