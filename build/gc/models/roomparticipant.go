package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoomparticipantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoomparticipantDud struct { 
    

}

// Roomparticipant
type Roomparticipant struct { 
    // ParticipantJid - participantJid
    ParticipantJid string `json:"participantJid"`

}

// String returns a JSON representation of the model
func (o *Roomparticipant) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Roomparticipant) MarshalJSON() ([]byte, error) {
    type Alias Roomparticipant

    if RoomparticipantMarshalled {
        return []byte("{}"), nil
    }
    RoomparticipantMarshalled = true

    return json.Marshal(&struct {
        
        ParticipantJid string `json:"participantJid"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

