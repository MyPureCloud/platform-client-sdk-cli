package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictedscoringMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictedscoringDud struct { }

// Predictedscoring
type Predictedscoring struct { }

// String returns a JSON representation of the model
func (o *Predictedscoring) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictedscoring) MarshalJSON() ([]byte, error) {
    type Alias Predictedscoring

    if PredictedscoringMarshalled {
        return []byte("{}"), nil
    }
    PredictedscoringMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

