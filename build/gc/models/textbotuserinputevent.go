package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotuserinputeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotuserinputeventDud struct { 
    


    

}

// Textbotuserinputevent - Settings for an input event to the bot flow indicating user input is available.
type Textbotuserinputevent struct { 
    // Mode - The input mode.
    Mode string `json:"mode"`


    // Alternatives - The input alternatives.
    Alternatives []Textbotuserinputalternative `json:"alternatives"`

}

// String returns a JSON representation of the model
func (o *Textbotuserinputevent) String() string {
    
     o.Alternatives = []Textbotuserinputalternative{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotuserinputevent) MarshalJSON() ([]byte, error) {
    type Alias Textbotuserinputevent

    if TextbotuserinputeventMarshalled {
        return []byte("{}"), nil
    }
    TextbotuserinputeventMarshalled = true

    return json.Marshal(&struct {
        
        Mode string `json:"mode"`
        
        Alternatives []Textbotuserinputalternative `json:"alternatives"`
        *Alias
    }{

        


        
        Alternatives: []Textbotuserinputalternative{{}},
        

        Alias: (*Alias)(u),
    })
}

