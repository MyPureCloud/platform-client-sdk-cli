package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediareactionschannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediareactionschannelDud struct { 
    Id string `json:"id"`


    Platform string `json:"platform"`


    VarType string `json:"type"`


    


    To Opensocialmediarecipient `json:"to"`


    

}

// Opensocialmediareactionschannel - Channel-specific information that describes the message and the message channel/provider.
type Opensocialmediareactionschannel struct { 
    


    


    


    // MessageId - Unique provider ID of the message such as a Open message ID.
    MessageId string `json:"messageId"`


    


    // Time - Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`

}

// String returns a JSON representation of the model
func (o *Opensocialmediareactionschannel) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediareactionschannel) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediareactionschannel

    if OpensocialmediareactionschannelMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediareactionschannelMarshalled = true

    return json.Marshal(&struct {
        
        MessageId string `json:"messageId"`
        
        Time time.Time `json:"time"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

