package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueuemessagingaddressesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueuemessagingaddressesDud struct { 
    

}

// Queuemessagingaddresses
type Queuemessagingaddresses struct { 
    // SmsAddress
    SmsAddress Domainentityref `json:"smsAddress"`

}

// String returns a JSON representation of the model
func (o *Queuemessagingaddresses) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queuemessagingaddresses) MarshalJSON() ([]byte, error) {
    type Alias Queuemessagingaddresses

    if QueuemessagingaddressesMarshalled {
        return []byte("{}"), nil
    }
    QueuemessagingaddressesMarshalled = true

    return json.Marshal(&struct { 
        SmsAddress Domainentityref `json:"smsAddress"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

