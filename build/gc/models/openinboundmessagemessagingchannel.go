package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundmessagemessagingchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundmessagemessagingchannelDud struct { 
    


    


    


    

}

// Openinboundmessagemessagingchannel - Open Channel-specific information that describes the message and the message channel/provider, with additional message information
type Openinboundmessagemessagingchannel struct { 
    // From - Information about the recipient the message is received from.
    From Openmessagingfromrecipient `json:"from"`


    // Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`


    // MessageId - Unique provider ID of the message.
    MessageId string `json:"messageId"`


    // Metadata - Additional Custom Information about the channel.
    Metadata Conversationchannelmetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Openinboundmessagemessagingchannel) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundmessagemessagingchannel) MarshalJSON() ([]byte, error) {
    type Alias Openinboundmessagemessagingchannel

    if OpeninboundmessagemessagingchannelMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundmessagemessagingchannelMarshalled = true

    return json.Marshal(&struct {
        
        From Openmessagingfromrecipient `json:"from"`
        
        Time time.Time `json:"time"`
        
        MessageId string `json:"messageId"`
        
        Metadata Conversationchannelmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

