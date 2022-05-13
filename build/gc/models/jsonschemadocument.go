package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JsonschemadocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JsonschemadocumentDud struct { 
    


    


    


    


    


    


    


    

}

// Jsonschemadocument - A JSON Schema document.
type Jsonschemadocument struct { 
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

}

// String returns a JSON representation of the model
func (o *Jsonschemadocument) String() string {
    
    
    
    
    
     o.Required = []string{""} 
     o.Properties = map[string]interface{}{"": Interface{}} 
     o.AdditionalProperties = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Jsonschemadocument) MarshalJSON() ([]byte, error) {
    type Alias Jsonschemadocument

    if JsonschemadocumentMarshalled {
        return []byte("{}"), nil
    }
    JsonschemadocumentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Schema string `json:"$schema"`
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        
        Required []string `json:"required"`
        
        Properties map[string]interface{} `json:"properties"`
        
        AdditionalProperties interface{} `json:"additionalProperties"`
        *Alias
    }{

        


        


        


        


        


        
        Required: []string{""},
        


        
        Properties: map[string]interface{}{"": Interface{}},
        


        
        AdditionalProperties: Interface{},
        

        Alias: (*Alias)(u),
    })
}

