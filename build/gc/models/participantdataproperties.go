package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParticipantdatapropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParticipantdatapropertiesDud struct { 
    

}

// Participantdataproperties
type Participantdataproperties struct { 
    // Name - Participant data name.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Participantdataproperties) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Participantdataproperties) MarshalJSON() ([]byte, error) {
    type Alias Participantdataproperties

    if ParticipantdatapropertiesMarshalled {
        return []byte("{}"), nil
    }
    ParticipantdatapropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

