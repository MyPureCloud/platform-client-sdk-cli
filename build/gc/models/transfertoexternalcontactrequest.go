package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransfertoexternalcontactrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransfertoexternalcontactrequestDud struct { 
    


    


    


    

}

// Transfertoexternalcontactrequest
type Transfertoexternalcontactrequest struct { 
    // TransferType - The type of transfer to perform. Attended, where the initiating agent maintains ownership of the conversation until the intended recipient accepts the transfer, or Unattended, where the initiating agent immediately disconnects. Default is Unattended.
    TransferType string `json:"transferType"`


    // KeepInternalMessageAlive - If true, the digital internal message will NOT be terminated.
    KeepInternalMessageAlive bool `json:"keepInternalMessageAlive"`


    // ContactId - The external contact id.
    ContactId string `json:"contactId"`


    // PhoneType - The external contact phone type.
    PhoneType string `json:"phoneType"`

}

// String returns a JSON representation of the model
func (o *Transfertoexternalcontactrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transfertoexternalcontactrequest) MarshalJSON() ([]byte, error) {
    type Alias Transfertoexternalcontactrequest

    if TransfertoexternalcontactrequestMarshalled {
        return []byte("{}"), nil
    }
    TransfertoexternalcontactrequestMarshalled = true

    return json.Marshal(&struct {
        
        TransferType string `json:"transferType"`
        
        KeepInternalMessageAlive bool `json:"keepInternalMessageAlive"`
        
        ContactId string `json:"contactId"`
        
        PhoneType string `json:"phoneType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

