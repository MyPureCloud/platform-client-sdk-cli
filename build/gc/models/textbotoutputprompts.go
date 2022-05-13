package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotoutputpromptsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotoutputpromptsDud struct { 
    


    

}

// Textbotoutputprompts - Prompt information related to a bot flow turn.
type Textbotoutputprompts struct { 
    // OutputLanguage - The ISO code of the output language for this prompt item.
    OutputLanguage string `json:"outputLanguage"`


    // TextPrompts - Text output prompts, if any.
    TextPrompts Textbotmodeoutputprompts `json:"textPrompts"`

}

// String returns a JSON representation of the model
func (o *Textbotoutputprompts) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotoutputprompts) MarshalJSON() ([]byte, error) {
    type Alias Textbotoutputprompts

    if TextbotoutputpromptsMarshalled {
        return []byte("{}"), nil
    }
    TextbotoutputpromptsMarshalled = true

    return json.Marshal(&struct {
        
        OutputLanguage string `json:"outputLanguage"`
        
        TextPrompts Textbotmodeoutputprompts `json:"textPrompts"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

