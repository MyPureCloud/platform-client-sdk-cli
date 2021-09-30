package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChannelentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChannelentitylistingDud struct { 
    

}

// Channelentitylisting
type Channelentitylisting struct { 
    // Entities
    Entities []Channel `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Channelentitylisting) String() string {
    
    
     o.Entities = []Channel{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Channelentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Channelentitylisting

    if ChannelentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ChannelentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Channel `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Channel{{}},
        

        
        Alias: (*Alias)(u),
    })
}

