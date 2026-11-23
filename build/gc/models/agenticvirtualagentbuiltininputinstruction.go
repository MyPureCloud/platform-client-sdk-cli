package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentbuiltininputinstructionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentbuiltininputinstructionDud struct { 
    


    


    


    

}

// Agenticvirtualagentbuiltininputinstruction - A built-in input instruction for a tool. Built-in instructions are stored separately from inputInstructions and merged into them at the given index when the version is published.
type Agenticvirtualagentbuiltininputinstruction struct { 
    // VarType - The kind of built-in instruction. A tool can have at most one built-in instruction of each type.
    VarType string `json:"type"`


    // Index - Position of this instruction in the tool's merged input instruction list. Must be unique per tool and less than the combined count of inputInstructions and builtinInputInstructions.
    Index int `json:"index"`


    // Override - Author-provided instruction text that replaces the default. Omit or set to null to use the default. Maximum 4000 characters.
    Override string `json:"override"`


    // VarDefault - The default instruction text. Used when no override is set. Maximum 4000 characters.
    VarDefault string `json:"default"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentbuiltininputinstruction) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentbuiltininputinstruction) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentbuiltininputinstruction

    if AgenticvirtualagentbuiltininputinstructionMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentbuiltininputinstructionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Index int `json:"index"`
        
        Override string `json:"override"`
        
        VarDefault string `json:"default"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

