package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChanneltopicentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChanneltopicentitylistingDud struct { 
    

}

// Channeltopicentitylisting
type Channeltopicentitylisting struct { 
    // Entities
    Entities []Channeltopic `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Channeltopicentitylisting) String() string {
    
    
     o.Entities = []Channeltopic{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Channeltopicentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Channeltopicentitylisting

    if ChanneltopicentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ChanneltopicentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Channeltopic `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Channeltopic{{}},
        

        
        Alias: (*Alias)(u),
    })
}

