package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FallbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FallbackDud struct { 
    


    

}

// Fallback
type Fallback struct { 
    // Enabled - Fallback actions are enabled.
    Enabled bool `json:"enabled"`


    // Actions - Fallback actions.
    Actions []Copilotfallbackaction `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Fallback) String() string {
    
     o.Actions = []Copilotfallbackaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fallback) MarshalJSON() ([]byte, error) {
    type Alias Fallback

    if FallbackMarshalled {
        return []byte("{}"), nil
    }
    FallbackMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        Actions []Copilotfallbackaction `json:"actions"`
        *Alias
    }{

        


        
        Actions: []Copilotfallbackaction{{}},
        

        Alias: (*Alias)(u),
    })
}

