package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UtteranceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UtteranceDud struct { 
    

}

// Utterance
type Utterance struct { 
    // UtteranceText - Utterance text
    UtteranceText string `json:"utteranceText"`

}

// String returns a JSON representation of the model
func (o *Utterance) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Utterance) MarshalJSON() ([]byte, error) {
    type Alias Utterance

    if UtteranceMarshalled {
        return []byte("{}"), nil
    }
    UtteranceMarshalled = true

    return json.Marshal(&struct {
        
        UtteranceText string `json:"utteranceText"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

