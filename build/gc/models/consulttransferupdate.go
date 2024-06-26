package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransferupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransferupdateDud struct { 
    


    

}

// Consulttransferupdate
type Consulttransferupdate struct { 
    // SpeakTo - Determines to whom the initiating participant is speaking.
    SpeakTo string `json:"speakTo"`


    // ConsultingUserId - The user ID of the person who wants to talk before completing the transfer. Could be the same of the context user ID
    ConsultingUserId string `json:"consultingUserId"`

}

// String returns a JSON representation of the model
func (o *Consulttransferupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransferupdate) MarshalJSON() ([]byte, error) {
    type Alias Consulttransferupdate

    if ConsulttransferupdateMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransferupdateMarshalled = true

    return json.Marshal(&struct {
        
        SpeakTo string `json:"speakTo"`
        
        ConsultingUserId string `json:"consultingUserId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

