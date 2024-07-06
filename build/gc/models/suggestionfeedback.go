package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionfeedbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionfeedbackDud struct { 
    

}

// Suggestionfeedback
type Suggestionfeedback struct { 
    // Rating - The rating value of the suggestion feedback.
    Rating string `json:"rating"`

}

// String returns a JSON representation of the model
func (o *Suggestionfeedback) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestionfeedback) MarshalJSON() ([]byte, error) {
    type Alias Suggestionfeedback

    if SuggestionfeedbackMarshalled {
        return []byte("{}"), nil
    }
    SuggestionfeedbackMarshalled = true

    return json.Marshal(&struct {
        
        Rating string `json:"rating"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

