package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParticipantmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParticipantmetricsDud struct { }

// Participantmetrics
type Participantmetrics struct { }

// String returns a JSON representation of the model
func (o *Participantmetrics) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Participantmetrics) MarshalJSON() ([]byte, error) {
    type Alias Participantmetrics

    if ParticipantmetricsMarshalled {
        return []byte("{}"), nil
    }
    ParticipantmetricsMarshalled = true

    return json.Marshal(&struct { 
        *Alias
    }{
        
        Alias: (*Alias)(u),
    })
}

