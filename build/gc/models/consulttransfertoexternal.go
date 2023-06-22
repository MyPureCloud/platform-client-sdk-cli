package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransfertoexternalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransfertoexternalDud struct { 
    


    


    

}

// Consulttransfertoexternal
type Consulttransfertoexternal struct { 
    // SpeakTo - Determines to whom the initiating participant is speaking. Defaults to DESTINATION
    SpeakTo string `json:"speakTo"`


    // ConsultingUserId - The user ID of the person who wants to talk before completing the transfer. Could be the same of the context user ID
    ConsultingUserId string `json:"consultingUserId"`


    // Address - The address (like phone number) of the external contact.
    Address string `json:"address"`

}

// String returns a JSON representation of the model
func (o *Consulttransfertoexternal) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransfertoexternal) MarshalJSON() ([]byte, error) {
    type Alias Consulttransfertoexternal

    if ConsulttransfertoexternalMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransfertoexternalMarshalled = true

    return json.Marshal(&struct {
        
        SpeakTo string `json:"speakTo"`
        
        ConsultingUserId string `json:"consultingUserId"`
        
        Address string `json:"address"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

