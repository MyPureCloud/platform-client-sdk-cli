package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenmessagingchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenmessagingchannelDud struct { 
    Id string `json:"id"`


    Platform string `json:"platform"`


    


    


    


    


    

}

// Openmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Openmessagingchannel struct { 
    


    


    // VarType - Specifies if this message is part of a private or public conversation.
    VarType string `json:"type"`


    // MessageId - Unique provider ID of the message such as a Facebook message ID.
    MessageId string `json:"messageId"`


    // To - Information about the recipient the message is sent to.
    To Openmessagingtorecipient `json:"to"`


    // From - Information about the recipient the message is received from.
    From Openmessagingfromrecipient `json:"from"`


    // Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`

}

// String returns a JSON representation of the model
func (o *Openmessagingchannel) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openmessagingchannel) MarshalJSON() ([]byte, error) {
    type Alias Openmessagingchannel

    if OpenmessagingchannelMarshalled {
        return []byte("{}"), nil
    }
    OpenmessagingchannelMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        VarType string `json:"type"`
        
        MessageId string `json:"messageId"`
        
        To Openmessagingtorecipient `json:"to"`
        
        From Openmessagingfromrecipient `json:"from"`
        
        Time time.Time `json:"time"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

