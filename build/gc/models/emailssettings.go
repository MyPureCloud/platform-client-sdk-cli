package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailssettingsDud struct { 
    

}

// Emailssettings
type Emailssettings struct { 
    // SendingSizeLimit
    SendingSizeLimit int `json:"sendingSizeLimit"`

}

// String returns a JSON representation of the model
func (o *Emailssettings) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailssettings) MarshalJSON() ([]byte, error) {
    type Alias Emailssettings

    if EmailssettingsMarshalled {
        return []byte("{}"), nil
    }
    EmailssettingsMarshalled = true

    return json.Marshal(&struct { 
        SendingSizeLimit int `json:"sendingSizeLimit"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

