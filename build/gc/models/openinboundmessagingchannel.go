package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundmessagingchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundmessagingchannelDud struct { 
    


    

}

// Openinboundmessagingchannel - Open Channel-specific information that describes the message and the message channel/provider.
type Openinboundmessagingchannel struct { 
    // From - Information about the recipient the message is received from.
    From Openmessagingfromrecipient `json:"from"`


    // Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`

}

// String returns a JSON representation of the model
func (o *Openinboundmessagingchannel) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundmessagingchannel) MarshalJSON() ([]byte, error) {
    type Alias Openinboundmessagingchannel

    if OpeninboundmessagingchannelMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundmessagingchannelMarshalled = true

    return json.Marshal(&struct {
        
        From Openmessagingfromrecipient `json:"from"`
        
        Time time.Time `json:"time"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

