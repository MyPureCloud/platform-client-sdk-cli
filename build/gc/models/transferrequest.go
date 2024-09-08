package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransferrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransferrequestDud struct { 
    


    


    


    


    


    


    

}

// Transferrequest
type Transferrequest struct { 
    // TransferType - The type of transfer to perform. Attended, where the initiating agent maintains ownership of the conversation until the intended recipient accepts the transfer, or Unattended, where the initiating agent immediately disconnects. Default is Unattended.
    TransferType string `json:"transferType"`


    // KeepInternalMessageAlive - If true, the digital internal message will NOT be terminated.
    KeepInternalMessageAlive bool `json:"keepInternalMessageAlive"`


    // UserId - The user ID of the transfer target.
    UserId string `json:"userId"`


    // Address - The user ID or queue ID of the transfer target. Address like a phone number can not be used for callbacks, but they can be used for other forms of communication.
    Address string `json:"address"`


    // UserName - The user name of the transfer target.
    UserName string `json:"userName"`


    // QueueId - The queue ID of the transfer target.
    QueueId string `json:"queueId"`


    // Voicemail - If true, transfer to the voicemail inbox of the participant that is being replaced.
    Voicemail bool `json:"voicemail"`

}

// String returns a JSON representation of the model
func (o *Transferrequest) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transferrequest) MarshalJSON() ([]byte, error) {
    type Alias Transferrequest

    if TransferrequestMarshalled {
        return []byte("{}"), nil
    }
    TransferrequestMarshalled = true

    return json.Marshal(&struct {
        
        TransferType string `json:"transferType"`
        
        KeepInternalMessageAlive bool `json:"keepInternalMessageAlive"`
        
        UserId string `json:"userId"`
        
        Address string `json:"address"`
        
        UserName string `json:"userName"`
        
        QueueId string `json:"queueId"`
        
        Voicemail bool `json:"voicemail"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

