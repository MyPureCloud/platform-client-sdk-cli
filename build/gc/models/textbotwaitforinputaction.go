package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotwaitforinputactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotwaitforinputactionDud struct { 
    

}

// Textbotwaitforinputaction - Settings for a next-action of waiting for additional user input and sending the data as an input action to the bot flow.
type Textbotwaitforinputaction struct { 
    // ModeConstraints - The mode constraints for the user input.
    ModeConstraints Textbotmodeconstraints `json:"modeConstraints"`

}

// String returns a JSON representation of the model
func (o *Textbotwaitforinputaction) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotwaitforinputaction) MarshalJSON() ([]byte, error) {
    type Alias Textbotwaitforinputaction

    if TextbotwaitforinputactionMarshalled {
        return []byte("{}"), nil
    }
    TextbotwaitforinputactionMarshalled = true

    return json.Marshal(&struct {
        
        ModeConstraints Textbotmodeconstraints `json:"modeConstraints"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

