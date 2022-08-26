package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationdisconnectsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationdisconnectsettingsDud struct { 
    


    

}

// Conversationdisconnectsettings
type Conversationdisconnectsettings struct { 
    // Enabled - whether or not conversation disconnect setting is enabled
    Enabled bool `json:"enabled"`


    // VarType - Conversation disconnect type
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Conversationdisconnectsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationdisconnectsettings) MarshalJSON() ([]byte, error) {
    type Alias Conversationdisconnectsettings

    if ConversationdisconnectsettingsMarshalled {
        return []byte("{}"), nil
    }
    ConversationdisconnectsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

