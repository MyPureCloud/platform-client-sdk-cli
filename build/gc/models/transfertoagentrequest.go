package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransfertoagentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransfertoagentrequestDud struct { 
    


    


    


    


    

}

// Transfertoagentrequest
type Transfertoagentrequest struct { 
    // TransferType - The type of transfer to perform. Attended, where the initiating agent maintains ownership of the conversation until the intended recipient accepts the transfer, or Unattended, where the initiating agent immediately disconnects. Default is Unattended.
    TransferType string `json:"transferType"`


    // UserId - The id of the internal user.
    UserId string `json:"userId"`


    // UserName - The userName (like user’s email) of the internal user.
    UserName string `json:"userName"`


    // UserDisplayName - The name of the internal user.
    UserDisplayName string `json:"userDisplayName"`


    // Voicemail - If true, transfer to the voicemail inbox of the participant that is being replaced.
    Voicemail bool `json:"voicemail"`

}

// String returns a JSON representation of the model
func (o *Transfertoagentrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transfertoagentrequest) MarshalJSON() ([]byte, error) {
    type Alias Transfertoagentrequest

    if TransfertoagentrequestMarshalled {
        return []byte("{}"), nil
    }
    TransfertoagentrequestMarshalled = true

    return json.Marshal(&struct {
        
        TransferType string `json:"transferType"`
        
        UserId string `json:"userId"`
        
        UserName string `json:"userName"`
        
        UserDisplayName string `json:"userDisplayName"`
        
        Voicemail bool `json:"voicemail"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

