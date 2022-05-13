package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotmodeconstraintsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotmodeconstraintsDud struct { 
    

}

// Textbotmodeconstraints - Mode constraints to observe when operating on a bot flow.
type Textbotmodeconstraints struct { 
    // Text - Mode constraints that apply to text scenarios.
    Text Textbottextmodeconstraints `json:"text"`

}

// String returns a JSON representation of the model
func (o *Textbotmodeconstraints) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotmodeconstraints) MarshalJSON() ([]byte, error) {
    type Alias Textbotmodeconstraints

    if TextbotmodeconstraintsMarshalled {
        return []byte("{}"), nil
    }
    TextbotmodeconstraintsMarshalled = true

    return json.Marshal(&struct {
        
        Text Textbottextmodeconstraints `json:"text"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

