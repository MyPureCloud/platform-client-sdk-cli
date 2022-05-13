package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotmodeoutputpromptsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotmodeoutputpromptsDud struct { 
    

}

// Textbotmodeoutputprompts - Prompt information related to a bot flow turn.
type Textbotmodeoutputprompts struct { 
    // Segments - The list of prompt segments.
    Segments []Textbotpromptsegment `json:"segments"`

}

// String returns a JSON representation of the model
func (o *Textbotmodeoutputprompts) String() string {
     o.Segments = []Textbotpromptsegment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotmodeoutputprompts) MarshalJSON() ([]byte, error) {
    type Alias Textbotmodeoutputprompts

    if TextbotmodeoutputpromptsMarshalled {
        return []byte("{}"), nil
    }
    TextbotmodeoutputpromptsMarshalled = true

    return json.Marshal(&struct {
        
        Segments []Textbotpromptsegment `json:"segments"`
        *Alias
    }{

        
        Segments: []Textbotpromptsegment{{}},
        

        Alias: (*Alias)(u),
    })
}

