package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AiscoringsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AiscoringsettingsDud struct { 
    


    

}

// Aiscoringsettings
type Aiscoringsettings struct { 
    // Enabled - True if AI scoring feature is configured, false if not.
    Enabled bool `json:"enabled"`


    // Prompt - The prompt text that would be used by a LLM.
    Prompt string `json:"prompt"`

}

// String returns a JSON representation of the model
func (o *Aiscoringsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aiscoringsettings) MarshalJSON() ([]byte, error) {
    type Alias Aiscoringsettings

    if AiscoringsettingsMarshalled {
        return []byte("{}"), nil
    }
    AiscoringsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        Prompt string `json:"prompt"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

