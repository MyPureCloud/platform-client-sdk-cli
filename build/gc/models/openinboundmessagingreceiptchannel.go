package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundmessagingreceiptchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundmessagingreceiptchannelDud struct { 
    


    

}

// Openinboundmessagingreceiptchannel - Open Channel-specific information that describes the message and the message channel/provider.
type Openinboundmessagingreceiptchannel struct { 
    // To - Information about the recipient the message is intended for.
    To Openmessagingtorecipient `json:"to"`


    // Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`

}

// String returns a JSON representation of the model
func (o *Openinboundmessagingreceiptchannel) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundmessagingreceiptchannel) MarshalJSON() ([]byte, error) {
    type Alias Openinboundmessagingreceiptchannel

    if OpeninboundmessagingreceiptchannelMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundmessagingreceiptchannelMarshalled = true

    return json.Marshal(&struct {
        
        To Openmessagingtorecipient `json:"to"`
        
        Time time.Time `json:"time"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

