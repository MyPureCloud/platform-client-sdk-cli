package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PresencesettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PresencesettingDud struct { 
    

}

// Presencesetting
type Presencesetting struct { 
    // Join - Should Presence Events be sent
    Join Settingdirection `json:"join"`

}

// String returns a JSON representation of the model
func (o *Presencesetting) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Presencesetting) MarshalJSON() ([]byte, error) {
    type Alias Presencesetting

    if PresencesettingMarshalled {
        return []byte("{}"), nil
    }
    PresencesettingMarshalled = true

    return json.Marshal(&struct { 
        Join Settingdirection `json:"join"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

