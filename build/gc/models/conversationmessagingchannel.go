package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagingchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagingchannelDud struct { 
    Id string `json:"id"`


    Platform string `json:"platform"`


    MessageId string `json:"messageId"`


    To Conversationmessagingtorecipient `json:"to"`


    From Conversationmessagingfromrecipient `json:"from"`


    Time time.Time `json:"time"`


    DateModified time.Time `json:"dateModified"`


    DateDeleted time.Time `json:"dateDeleted"`

}

// Conversationmessagingchannel - Channel-specific information that describes the message and the message channel/provider.
type Conversationmessagingchannel struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Conversationmessagingchannel) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagingchannel) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagingchannel

    if ConversationmessagingchannelMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagingchannelMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

