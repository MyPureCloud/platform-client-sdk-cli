package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MaxparticipantsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MaxparticipantsDud struct { 
    

}

// Maxparticipants
type Maxparticipants struct { 
    // MaxParticipants - The maximum number of participants that are allowed on a conversation.
    MaxParticipants int `json:"maxParticipants"`

}

// String returns a JSON representation of the model
func (o *Maxparticipants) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Maxparticipants) MarshalJSON() ([]byte, error) {
    type Alias Maxparticipants

    if MaxparticipantsMarshalled {
        return []byte("{}"), nil
    }
    MaxparticipantsMarshalled = true

    return json.Marshal(&struct { 
        MaxParticipants int `json:"maxParticipants"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

