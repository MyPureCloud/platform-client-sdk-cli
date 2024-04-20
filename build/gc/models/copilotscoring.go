package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotscoringMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotscoringDud struct { 
    


    

}

// Copilotscoring
type Copilotscoring struct { 
    // Enabled - True if copilot feature is configured, false if not.
    Enabled bool `json:"enabled"`


    // Prompt - The prompt text that would be used by a LLM.
    Prompt string `json:"prompt"`

}

// String returns a JSON representation of the model
func (o *Copilotscoring) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotscoring) MarshalJSON() ([]byte, error) {
    type Alias Copilotscoring

    if CopilotscoringMarshalled {
        return []byte("{}"), nil
    }
    CopilotscoringMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        Prompt string `json:"prompt"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

