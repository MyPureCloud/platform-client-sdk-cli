package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HumanizeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HumanizeDud struct { 
    


    

}

// Humanize
type Humanize struct { 
    // Enabled - Whether or not humanize conversations setting is enabled
    Enabled bool `json:"enabled"`


    // Bot - Bot messenger profile setting
    Bot Botmessengerprofile `json:"bot"`

}

// String returns a JSON representation of the model
func (o *Humanize) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Humanize) MarshalJSON() ([]byte, error) {
    type Alias Humanize

    if HumanizeMarshalled {
        return []byte("{}"), nil
    }
    HumanizeMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        Bot Botmessengerprofile `json:"bot"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

