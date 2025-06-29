package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransfertoaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransfertoaddressDud struct { 
    


    


    

}

// Consulttransfertoaddress
type Consulttransfertoaddress struct { 
    // SpeakTo - Determines to whom the initiating participant is requesting to speak. Defaults to DESTINATION
    SpeakTo string `json:"speakTo"`


    // ConsultingUserId - The user ID of the person who wants to talk before completing the transfer. Could be the same of the context user ID
    ConsultingUserId string `json:"consultingUserId"`


    // Address - Destination's address or phone number.
    Address string `json:"address"`

}

// String returns a JSON representation of the model
func (o *Consulttransfertoaddress) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransfertoaddress) MarshalJSON() ([]byte, error) {
    type Alias Consulttransfertoaddress

    if ConsulttransfertoaddressMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransfertoaddressMarshalled = true

    return json.Marshal(&struct {
        
        SpeakTo string `json:"speakTo"`
        
        ConsultingUserId string `json:"consultingUserId"`
        
        Address string `json:"address"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

