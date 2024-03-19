package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoomparticipantresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoomparticipantresponseDud struct { 
    


    

}

// Roomparticipantresponse
type Roomparticipantresponse struct { 
    // Jid - jid of the participant
    Jid string `json:"jid"`


    // User - User id of the participant
    User Addressableentityref `json:"user"`

}

// String returns a JSON representation of the model
func (o *Roomparticipantresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Roomparticipantresponse) MarshalJSON() ([]byte, error) {
    type Alias Roomparticipantresponse

    if RoomparticipantresponseMarshalled {
        return []byte("{}"), nil
    }
    RoomparticipantresponseMarshalled = true

    return json.Marshal(&struct {
        
        Jid string `json:"jid"`
        
        User Addressableentityref `json:"user"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

