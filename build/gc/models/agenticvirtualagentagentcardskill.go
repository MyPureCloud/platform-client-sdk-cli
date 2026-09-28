package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentagentcardskillMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentagentcardskillDud struct { 
    


    


    


    


    


    


    

}

// Agenticvirtualagentagentcardskill - A2A agent card skill.
type Agenticvirtualagentagentcardskill struct { 
    // Id - Unique identifier for the skill.
    Id string `json:"id"`


    // Name - Human-readable name of the skill.
    Name string `json:"name"`


    // Description - Detailed explanation of what the skill does.
    Description string `json:"description"`


    // Tags - Keywords for categorization and discovery.
    Tags []string `json:"tags"`


    // Examples - Sample prompts or use cases.
    Examples []string `json:"examples"`


    // InputModes - Supported input media types.
    InputModes []string `json:"inputModes"`


    // OutputModes - Supported output media types.
    OutputModes []string `json:"outputModes"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentagentcardskill) String() string {
    
    
    
     o.Tags = []string{""} 
     o.Examples = []string{""} 
     o.InputModes = []string{""} 
     o.OutputModes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentagentcardskill) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentagentcardskill

    if AgenticvirtualagentagentcardskillMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentagentcardskillMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Tags []string `json:"tags"`
        
        Examples []string `json:"examples"`
        
        InputModes []string `json:"inputModes"`
        
        OutputModes []string `json:"outputModes"`
        *Alias
    }{

        


        


        


        
        Tags: []string{""},
        


        
        Examples: []string{""},
        


        
        InputModes: []string{""},
        


        
        OutputModes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

