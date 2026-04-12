package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntakesettingsupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntakesettingsupdateDud struct { 
    

}

// Intakesettingsupdate
type Intakesettingsupdate struct { 
    // IntakeSettings - The intake settings for the Caseplan.
    IntakeSettings []Intakesetting `json:"intakeSettings"`

}

// String returns a JSON representation of the model
func (o *Intakesettingsupdate) String() string {
     o.IntakeSettings = []Intakesetting{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intakesettingsupdate) MarshalJSON() ([]byte, error) {
    type Alias Intakesettingsupdate

    if IntakesettingsupdateMarshalled {
        return []byte("{}"), nil
    }
    IntakesettingsupdateMarshalled = true

    return json.Marshal(&struct {
        
        IntakeSettings []Intakesetting `json:"intakeSettings"`
        *Alias
    }{

        
        IntakeSettings: []Intakesetting{{}},
        

        Alias: (*Alias)(u),
    })
}

