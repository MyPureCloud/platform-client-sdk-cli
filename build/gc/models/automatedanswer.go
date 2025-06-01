package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AutomatedanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AutomatedanswerDud struct { }

// Automatedanswer
type Automatedanswer struct { }

// String returns a JSON representation of the model
func (o *Automatedanswer) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Automatedanswer) MarshalJSON() ([]byte, error) {
    type Alias Automatedanswer

    if AutomatedanswerMarshalled {
        return []byte("{}"), nil
    }
    AutomatedanswerMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

