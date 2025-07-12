package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutbounddomainrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutbounddomainrequestDud struct { 
    


    


    

}

// Outbounddomainrequest
type Outbounddomainrequest struct { 
    // Id - Unique Id of the domain such as: example.com
    Id string `json:"id"`


    // SenderType - Sender Type
    SenderType string `json:"senderType"`


    // Name - The domain such as: example.com
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Outbounddomainrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outbounddomainrequest) MarshalJSON() ([]byte, error) {
    type Alias Outbounddomainrequest

    if OutbounddomainrequestMarshalled {
        return []byte("{}"), nil
    }
    OutbounddomainrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SenderType string `json:"senderType"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

