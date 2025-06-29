package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransfertoaddressrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransfertoaddressrequestDud struct { 
    


    


    

}

// Transfertoaddressrequest
type Transfertoaddressrequest struct { 
    // TransferType - The type of transfer to perform. Attended, where the initiating agent maintains ownership of the conversation until the intended recipient accepts the transfer, or Unattended, where the initiating agent immediately disconnects. Default is Unattended.
    TransferType string `json:"transferType"`


    // KeepInternalMessageAlive - If true, the digital internal message will NOT be terminated.
    KeepInternalMessageAlive bool `json:"keepInternalMessageAlive"`


    // Address - User's name, queue's name, destination's address or phone number.
    Address string `json:"address"`

}

// String returns a JSON representation of the model
func (o *Transfertoaddressrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transfertoaddressrequest) MarshalJSON() ([]byte, error) {
    type Alias Transfertoaddressrequest

    if TransfertoaddressrequestMarshalled {
        return []byte("{}"), nil
    }
    TransfertoaddressrequestMarshalled = true

    return json.Marshal(&struct {
        
        TransferType string `json:"transferType"`
        
        KeepInternalMessageAlive bool `json:"keepInternalMessageAlive"`
        
        Address string `json:"address"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

