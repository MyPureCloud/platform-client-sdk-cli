package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AcwsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AcwsettingsDud struct { 
    


    

}

// Acwsettings
type Acwsettings struct { 
    // WrapupPrompt - This field controls how the UI prompts the agent for a wrapup.
    WrapupPrompt string `json:"wrapupPrompt"`


    // TimeoutMs - The amount of time the agent can stay in ACW (Min: 1 sec, Max: 60 min).  Can only be used when ACW is MANDATORY_TIMEOUT or MANDATORY_FORCED_TIMEOUT.
    TimeoutMs int `json:"timeoutMs"`

}

// String returns a JSON representation of the model
func (o *Acwsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Acwsettings) MarshalJSON() ([]byte, error) {
    type Alias Acwsettings

    if AcwsettingsMarshalled {
        return []byte("{}"), nil
    }
    AcwsettingsMarshalled = true

    return json.Marshal(&struct {
        
        WrapupPrompt string `json:"wrapupPrompt"`
        
        TimeoutMs int `json:"timeoutMs"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

