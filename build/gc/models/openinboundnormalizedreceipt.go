package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundnormalizedreceiptMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundnormalizedreceiptDud struct { 
    


    


    


    


    

}

// Openinboundnormalizedreceipt - Open Messaging rich media message structure
type Openinboundnormalizedreceipt struct { 
    // Id - The original unique message Id generated by the messaging platform, that this receipt message is referencing.
    Id string `json:"id"`


    // Channel - Channel-specific information that describes the message and the message channel/provider.
    Channel Openinboundmessagingreceiptchannel `json:"channel"`


    // Status - Message receipt status.
    Status string `json:"status"`


    // Reasons - List of reasons for a message receipt that indicates the message has failed. Only used with Failed status.
    Reasons []Conversationreason `json:"reasons"`


    // IsFinalReceipt - Indicates if this is the last message receipt for this message, or if another message receipt can be expected.
    IsFinalReceipt bool `json:"isFinalReceipt"`

}

// String returns a JSON representation of the model
func (o *Openinboundnormalizedreceipt) String() string {
    
    
    
     o.Reasons = []Conversationreason{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundnormalizedreceipt) MarshalJSON() ([]byte, error) {
    type Alias Openinboundnormalizedreceipt

    if OpeninboundnormalizedreceiptMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundnormalizedreceiptMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Channel Openinboundmessagingreceiptchannel `json:"channel"`
        
        Status string `json:"status"`
        
        Reasons []Conversationreason `json:"reasons"`
        
        IsFinalReceipt bool `json:"isFinalReceipt"`
        *Alias
    }{

        


        


        


        
        Reasons: []Conversationreason{{}},
        


        

        Alias: (*Alias)(u),
    })
}
