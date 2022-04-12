package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TypingsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TypingsettingDud struct { 
    

}

// Typingsetting
type Typingsetting struct { 
    // On - Should typing indication Events be sent
    On Settingdirection `json:"on"`

}

// String returns a JSON representation of the model
func (o *Typingsetting) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Typingsetting) MarshalJSON() ([]byte, error) {
    type Alias Typingsetting

    if TypingsettingMarshalled {
        return []byte("{}"), nil
    }
    TypingsettingMarshalled = true

    return json.Marshal(&struct { 
        On Settingdirection `json:"on"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

