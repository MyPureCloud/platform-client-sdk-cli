package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingmessagereceiptMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingmessagereceiptDud struct { 
    


    


    

}

// Recordingmessagereceipt
type Recordingmessagereceipt struct { 
    // Id - The id of the message receipt. Message receipts will have the same ID as the message they reference.
    Id string `json:"id"`


    // Status - The message receipt status
    Status string `json:"status"`


    // Reasons - List of reasons for a message receipt that indicates the message has failed. Only used with Failed status.
    Reasons []Recordingmessagereceiptreason `json:"reasons"`

}

// String returns a JSON representation of the model
func (o *Recordingmessagereceipt) String() string {
    
    
     o.Reasons = []Recordingmessagereceiptreason{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingmessagereceipt) MarshalJSON() ([]byte, error) {
    type Alias Recordingmessagereceipt

    if RecordingmessagereceiptMarshalled {
        return []byte("{}"), nil
    }
    RecordingmessagereceiptMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        Reasons []Recordingmessagereceiptreason `json:"reasons"`
        *Alias
    }{

        


        


        
        Reasons: []Recordingmessagereceiptreason{{}},
        

        Alias: (*Alias)(u),
    })
}

