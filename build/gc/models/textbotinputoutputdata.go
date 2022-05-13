package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotinputoutputdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotinputoutputdataDud struct { 
    

}

// Textbotinputoutputdata - Input/Output data related to a bot flow which is exiting gracefully.
type Textbotinputoutputdata struct { 
    // Variables - The input/output variables using the format as appropriate for the variable data type in the flow definition.
    Variables map[string]interface{} `json:"variables"`

}

// String returns a JSON representation of the model
func (o *Textbotinputoutputdata) String() string {
     o.Variables = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotinputoutputdata) MarshalJSON() ([]byte, error) {
    type Alias Textbotinputoutputdata

    if TextbotinputoutputdataMarshalled {
        return []byte("{}"), nil
    }
    TextbotinputoutputdataMarshalled = true

    return json.Marshal(&struct {
        
        Variables map[string]interface{} `json:"variables"`
        *Alias
    }{

        
        Variables: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

