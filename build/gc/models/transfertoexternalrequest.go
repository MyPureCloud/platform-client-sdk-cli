package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransfertoexternalrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransfertoexternalrequestDud struct { 
    


    

}

// Transfertoexternalrequest
type Transfertoexternalrequest struct { 
    // TransferType
    TransferType string `json:"transferType"`


    // Address - The address (like phone number) of the external contact.
    Address string `json:"address"`

}

// String returns a JSON representation of the model
func (o *Transfertoexternalrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transfertoexternalrequest) MarshalJSON() ([]byte, error) {
    type Alias Transfertoexternalrequest

    if TransfertoexternalrequestMarshalled {
        return []byte("{}"), nil
    }
    TransfertoexternalrequestMarshalled = true

    return json.Marshal(&struct {
        
        TransferType string `json:"transferType"`
        
        Address string `json:"address"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

