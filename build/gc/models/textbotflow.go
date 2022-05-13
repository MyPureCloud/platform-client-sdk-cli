package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotflowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotflowDud struct { 
    

}

// Textbotflow - Description of the Bot Flow.
type Textbotflow struct { 
    // Id - The Bot Flow ID.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Textbotflow) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotflow) MarshalJSON() ([]byte, error) {
    type Alias Textbotflow

    if TextbotflowMarshalled {
        return []byte("{}"), nil
    }
    TextbotflowMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

