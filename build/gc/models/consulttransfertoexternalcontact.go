package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransfertoexternalcontactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransfertoexternalcontactDud struct { 
    


    


    


    

}

// Consulttransfertoexternalcontact
type Consulttransfertoexternalcontact struct { 
    // SpeakTo - Determines to whom the initiating participant is requesting to speak. Defaults to DESTINATION
    SpeakTo string `json:"speakTo"`


    // ConsultingUserId - The user ID of the person who wants to talk before completing the transfer. Could be the same of the context user ID
    ConsultingUserId string `json:"consultingUserId"`


    // ContactId - The external contact id.
    ContactId string `json:"contactId"`


    // PhoneType - The external contact phone type.
    PhoneType string `json:"phoneType"`

}

// String returns a JSON representation of the model
func (o *Consulttransfertoexternalcontact) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransfertoexternalcontact) MarshalJSON() ([]byte, error) {
    type Alias Consulttransfertoexternalcontact

    if ConsulttransfertoexternalcontactMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransfertoexternalcontactMarshalled = true

    return json.Marshal(&struct {
        
        SpeakTo string `json:"speakTo"`
        
        ConsultingUserId string `json:"consultingUserId"`
        
        ContactId string `json:"contactId"`
        
        PhoneType string `json:"phoneType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

