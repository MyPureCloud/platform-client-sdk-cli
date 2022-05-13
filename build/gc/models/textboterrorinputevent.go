package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextboterrorinputeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextboterrorinputeventDud struct { 
    


    

}

// Textboterrorinputevent - Settings for an input event to the bot flow indicating an error has occurred.
type Textboterrorinputevent struct { 
    // Code - The error code.
    Code string `json:"code"`


    // Message - The error message.
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Textboterrorinputevent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textboterrorinputevent) MarshalJSON() ([]byte, error) {
    type Alias Textboterrorinputevent

    if TextboterrorinputeventMarshalled {
        return []byte("{}"), nil
    }
    TextboterrorinputeventMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

