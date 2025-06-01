package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediachannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediachannelDud struct { 
    Id string `json:"id"`


    Platform string `json:"platform"`


    VarType string `json:"type"`


    


    


    


    


    


    

}

// Opensocialmediachannel - Channel-specific information that describes the message and the message channel/provider.
type Opensocialmediachannel struct { 
    


    


    


    // MessageId - Unique provider ID of the message such as a Facebook message ID.
    MessageId string `json:"messageId"`


    // To - Information about the recipient the message is sent to.
    To Opensocialmediarecipient `json:"to"`


    // From - Information about the recipient the message is received from.
    From Opensocialmediarecipient `json:"from"`


    // Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`


    // Metadata - Information about the channel.
    Metadata interface{} `json:"metadata"`


    // PublicMetadata - Meta data of this public post. For example, used to define where in the thread the post exists.
    PublicMetadata Opensocialmediapublicmetadata `json:"publicMetadata"`

}

// String returns a JSON representation of the model
func (o *Opensocialmediachannel) String() string {
    
    
    
    
     o.Metadata = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediachannel) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediachannel

    if OpensocialmediachannelMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediachannelMarshalled = true

    return json.Marshal(&struct {
        
        MessageId string `json:"messageId"`
        
        To Opensocialmediarecipient `json:"to"`
        
        From Opensocialmediarecipient `json:"from"`
        
        Time time.Time `json:"time"`
        
        Metadata interface{} `json:"metadata"`
        
        PublicMetadata Opensocialmediapublicmetadata `json:"publicMetadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Metadata: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}

