package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagenttypedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagenttypedefinitionDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Agenticvirtualagenttypedefinition - Type definition used by a virtual agent. The applicable fields depend on the type value and related fields.
type Agenticvirtualagenttypedefinition struct { 
    // Name - Type name.
    Name string `json:"name"`


    // Description - Additional context that helps the virtual agent understand what this type is used for.
    Description string `json:"description"`


    // Direction - Intended direction of use for this type.
    Direction string `json:"direction"`


    // VarType - Type value. The applicable fields depend on this value and related fields.
    VarType string `json:"type"`


    // UserUtteranceSubstring - Whether values of this string type must be copied as a contiguous substring from recent user messages.
    UserUtteranceSubstring bool `json:"userUtteranceSubstring"`


    // Undisclosed - Whether values of this string type are hidden from the virtual agent and represented as opaque identifiers. Only valid when type is string.
    Undisclosed bool `json:"undisclosed"`


    // Properties - Properties of this object type. Applies when type is object.
    Properties []Agenticvirtualagentpropertydefinition `json:"properties"`


    // Items - Type of items in this array type. Applies when type is array.
    Items string `json:"items"`


    // StatusCodes - HTTP 4xx or 5xx status codes this error type can handle. Applies when type is DataActionHttpError.
    StatusCodes []int `json:"statusCodes"`


    // DefaultInstruction - Default instruction for how the virtual agent should handle this error type when a tool references it without its own error instruction. Applies when type is DataActionHttpError.
    DefaultInstruction string `json:"defaultInstruction"`


    // Enum - Allowed enum values. Applies to enum types.
    Enum []string `json:"enum"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttypedefinition) String() string {
    
    
    
    
    
    
     o.Properties = []Agenticvirtualagentpropertydefinition{{}} 
    
     o.StatusCodes = []int{0} 
    
     o.Enum = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagenttypedefinition) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagenttypedefinition

    if AgenticvirtualagenttypedefinitionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagenttypedefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Direction string `json:"direction"`
        
        VarType string `json:"type"`
        
        UserUtteranceSubstring bool `json:"userUtteranceSubstring"`
        
        Undisclosed bool `json:"undisclosed"`
        
        Properties []Agenticvirtualagentpropertydefinition `json:"properties"`
        
        Items string `json:"items"`
        
        StatusCodes []int `json:"statusCodes"`
        
        DefaultInstruction string `json:"defaultInstruction"`
        
        Enum []string `json:"enum"`
        *Alias
    }{

        


        


        


        


        


        


        
        Properties: []Agenticvirtualagentpropertydefinition{{}},
        


        


        
        StatusCodes: []int{0},
        


        


        
        Enum: []string{""},
        

        Alias: (*Alias)(u),
    })
}

