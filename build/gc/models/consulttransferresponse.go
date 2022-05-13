package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransferresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransferresponseDud struct { 
    

}

// Consulttransferresponse
type Consulttransferresponse struct { 
    // DestinationParticipantId - Participant ID to whom the call is being transferred.
    DestinationParticipantId string `json:"destinationParticipantId"`

}

// String returns a JSON representation of the model
func (o *Consulttransferresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransferresponse) MarshalJSON() ([]byte, error) {
    type Alias Consulttransferresponse

    if ConsulttransferresponseMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransferresponseMarshalled = true

    return json.Marshal(&struct {
        
        DestinationParticipantId string `json:"destinationParticipantId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

