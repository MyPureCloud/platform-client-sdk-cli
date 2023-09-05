package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundnormalizedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundnormalizedeventDud struct { 
    


    

}

// Openinboundnormalizedevent - Open Event Messaging rich media message structure
type Openinboundnormalizedevent struct { 
    // Channel - Channel-specific information that describes the message and the message channel/provider.
    Channel Openinboundmessagingchannel `json:"channel"`


    // Events - List of event elements.
    Events []Openevent `json:"events"`

}

// String returns a JSON representation of the model
func (o *Openinboundnormalizedevent) String() string {
    
     o.Events = []Openevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundnormalizedevent) MarshalJSON() ([]byte, error) {
    type Alias Openinboundnormalizedevent

    if OpeninboundnormalizedeventMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundnormalizedeventMarshalled = true

    return json.Marshal(&struct {
        
        Channel Openinboundmessagingchannel `json:"channel"`
        
        Events []Openevent `json:"events"`
        *Alias
    }{

        


        
        Events: []Openevent{{}},
        

        Alias: (*Alias)(u),
    })
}

