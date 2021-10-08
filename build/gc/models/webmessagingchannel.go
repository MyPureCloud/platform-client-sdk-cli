package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingchannelDud struct { 
    From Webmessagingrecipient `json:"from"`


    To Webmessagingrecipient `json:"to"`


    Time time.Time `json:"time"`


    MessageId string `json:"messageId"`

}

// Webmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Webmessagingchannel struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Webmessagingchannel) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingchannel) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingchannel

    if WebmessagingchannelMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingchannelMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

