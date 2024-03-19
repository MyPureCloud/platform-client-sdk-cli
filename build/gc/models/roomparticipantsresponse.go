package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoomparticipantsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoomparticipantsresponseDud struct { 
    

}

// Roomparticipantsresponse
type Roomparticipantsresponse struct { 
    // Participants - list of room participants
    Participants []Roomparticipantresponse `json:"participants"`

}

// String returns a JSON representation of the model
func (o *Roomparticipantsresponse) String() string {
     o.Participants = []Roomparticipantresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Roomparticipantsresponse) MarshalJSON() ([]byte, error) {
    type Alias Roomparticipantsresponse

    if RoomparticipantsresponseMarshalled {
        return []byte("{}"), nil
    }
    RoomparticipantsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Participants []Roomparticipantresponse `json:"participants"`
        *Alias
    }{

        
        Participants: []Roomparticipantresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

