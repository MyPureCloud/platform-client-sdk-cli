package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentrepetitioncheckMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentrepetitioncheckDud struct { 
    


    


    


    

}

// Agenticvirtualagentrepetitioncheck - A rule that detects repeated user or agent messages and adds a corrective instruction.
type Agenticvirtualagentrepetitioncheck struct { 
    // VarType - Whether this check looks for repetition in user messages or agent responses.
    VarType string `json:"type"`


    // Messages - The number of prior messages of the specified type to compare for repetition.
    Messages int `json:"messages"`


    // Similarity - The similarity category compared to the Levenshtein result that triggers this check's instruction.
    Similarity string `json:"similarity"`


    // Instruction - The instruction added to the virtual agent's turn when message similarity matches the configured category.
    Instruction string `json:"instruction"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentrepetitioncheck) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentrepetitioncheck) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentrepetitioncheck

    if AgenticvirtualagentrepetitioncheckMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentrepetitioncheckMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Messages int `json:"messages"`
        
        Similarity string `json:"similarity"`
        
        Instruction string `json:"instruction"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

