package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransfertoagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransfertoagentDud struct { 
    


    


    


    

}

// Consulttransfertoagent
type Consulttransfertoagent struct { 
    // SpeakTo - Determines to whom the initiating participant is requesting to speak. Defaults to DESTINATION
    SpeakTo string `json:"speakTo"`


    // ConsultingUserId - The user ID of the person who wants to talk before completing the transfer. Could be the same of the context user ID
    ConsultingUserId string `json:"consultingUserId"`


    // UserId - The id of the internal user.
    UserId string `json:"userId"`


    // UserDisplayName - The name of the internal user.
    UserDisplayName string `json:"userDisplayName"`

}

// String returns a JSON representation of the model
func (o *Consulttransfertoagent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransfertoagent) MarshalJSON() ([]byte, error) {
    type Alias Consulttransfertoagent

    if ConsulttransfertoagentMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransfertoagentMarshalled = true

    return json.Marshal(&struct {
        
        SpeakTo string `json:"speakTo"`
        
        ConsultingUserId string `json:"consultingUserId"`
        
        UserId string `json:"userId"`
        
        UserDisplayName string `json:"userDisplayName"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

