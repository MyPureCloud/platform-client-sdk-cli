package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentpropertydefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentpropertydefinitionDud struct { 
    


    


    


    


    


    

}

// Agenticvirtualagentpropertydefinition - Property definition for an object type.
type Agenticvirtualagentpropertydefinition struct { 
    // Name - Property name.
    Name string `json:"name"`


    // VarType - Property type name. The valid type depends on the containing type and related fields.
    VarType string `json:"type"`


    // Required - Whether this property must be supplied.
    Required bool `json:"required"`


    // Description - Additional context that helps the virtual agent understand what this property means.
    Description string `json:"description"`


    // Items - Type of items in this array property. Applies when type is array.
    Items string `json:"items"`


    // Mapping - Path used to extract this output data property from a tool output. Only valid for output data properties. The path starts with a tool output type name, may contain only string property names or integer array indexes, and must resolve to a primitive value.
    Mapping []interface{} `json:"mapping"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentpropertydefinition) String() string {
    
    
    
    
    
     o.Mapping = []interface{}{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentpropertydefinition) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentpropertydefinition

    if AgenticvirtualagentpropertydefinitionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentpropertydefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Required bool `json:"required"`
        
        Description string `json:"description"`
        
        Items string `json:"items"`
        
        Mapping []interface{} `json:"mapping"`
        *Alias
    }{

        


        


        


        


        


        
        Mapping: []interface{}{},
        

        Alias: (*Alias)(u),
    })
}

