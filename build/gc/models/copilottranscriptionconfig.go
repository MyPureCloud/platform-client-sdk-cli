package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilottranscriptionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilottranscriptionconfigDud struct { 
    

}

// Copilottranscriptionconfig
type Copilottranscriptionconfig struct { 
    // Engine - The Transcription engine for Agent Copilot.
    Engine string `json:"engine"`

}

// String returns a JSON representation of the model
func (o *Copilottranscriptionconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilottranscriptionconfig) MarshalJSON() ([]byte, error) {
    type Alias Copilottranscriptionconfig

    if CopilottranscriptionconfigMarshalled {
        return []byte("{}"), nil
    }
    CopilottranscriptionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Engine string `json:"engine"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

