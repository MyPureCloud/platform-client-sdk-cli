package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediareactionsnormalizedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediareactionsnormalizedeventDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Opensocialmediareactionsnormalizedevent - Open Social Messaging rich media event structure
type Opensocialmediareactionsnormalizedevent struct { 
    


    // Channel - Channel-specific information that describes the message and the message channel/provider.
    Channel Opensocialmediareactionschannel `json:"channel"`


    // Events - List of event elements.
    Events []Opensocialmediareactionsevent `json:"events"`


    

}

// String returns a JSON representation of the model
func (o *Opensocialmediareactionsnormalizedevent) String() string {
    
     o.Events = []Opensocialmediareactionsevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediareactionsnormalizedevent) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediareactionsnormalizedevent

    if OpensocialmediareactionsnormalizedeventMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediareactionsnormalizedeventMarshalled = true

    return json.Marshal(&struct {
        
        Channel Opensocialmediareactionschannel `json:"channel"`
        
        Events []Opensocialmediareactionsevent `json:"events"`
        *Alias
    }{

        


        


        
        Events: []Opensocialmediareactionsevent{{}},
        


        

        Alias: (*Alias)(u),
    })
}

