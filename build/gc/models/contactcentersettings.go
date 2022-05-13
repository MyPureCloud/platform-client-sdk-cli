package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactcentersettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactcentersettingsDud struct { 
    

}

// Contactcentersettings
type Contactcentersettings struct { 
    // RemoveSkillsFromBlindTransfer - Strip skills from transfer
    RemoveSkillsFromBlindTransfer bool `json:"removeSkillsFromBlindTransfer"`

}

// String returns a JSON representation of the model
func (o *Contactcentersettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactcentersettings) MarshalJSON() ([]byte, error) {
    type Alias Contactcentersettings

    if ContactcentersettingsMarshalled {
        return []byte("{}"), nil
    }
    ContactcentersettingsMarshalled = true

    return json.Marshal(&struct {
        
        RemoveSkillsFromBlindTransfer bool `json:"removeSkillsFromBlindTransfer"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

