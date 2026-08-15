package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentdynamicturninstructionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentdynamicturninstructionsDud struct { 
    

}

// Agenticvirtualagentdynamicturninstructions - Instructions dynamically added to the virtual agent based on conversation state.
type Agenticvirtualagentdynamicturninstructions struct { 
    // RepetitionChecks - Checks that can be configured to add dynamic instructions for the agent, if user / agent messages repeat.
    RepetitionChecks []Agenticvirtualagentrepetitioncheck `json:"repetitionChecks"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentdynamicturninstructions) String() string {
     o.RepetitionChecks = []Agenticvirtualagentrepetitioncheck{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentdynamicturninstructions) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentdynamicturninstructions

    if AgenticvirtualagentdynamicturninstructionsMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentdynamicturninstructionsMarshalled = true

    return json.Marshal(&struct {
        
        RepetitionChecks []Agenticvirtualagentrepetitioncheck `json:"repetitionChecks"`
        *Alias
    }{

        
        RepetitionChecks: []Agenticvirtualagentrepetitioncheck{{}},
        

        Alias: (*Alias)(u),
    })
}

