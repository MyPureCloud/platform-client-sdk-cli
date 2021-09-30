package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediatypeaccessMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediatypeaccessDud struct { 
    


    

}

// Mediatypeaccess - Media type access definitions
type Mediatypeaccess struct { 
    // Inbound - List of media types allowed for inbound messages from customers. If inbound messages from a customer contain media that is not in this list, the media will be dropped from the outbound message.
    Inbound []Mediatype `json:"inbound"`


    // Outbound - List of media types allowed for outbound messages to customers. If an outbound message is sent that contains media that is not in this list, the message will not be sent.
    Outbound []Mediatype `json:"outbound"`

}

// String returns a JSON representation of the model
func (o *Mediatypeaccess) String() string {
    
    
     o.Inbound = []Mediatype{{}} 
    
    
    
     o.Outbound = []Mediatype{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediatypeaccess) MarshalJSON() ([]byte, error) {
    type Alias Mediatypeaccess

    if MediatypeaccessMarshalled {
        return []byte("{}"), nil
    }
    MediatypeaccessMarshalled = true

    return json.Marshal(&struct { 
        Inbound []Mediatype `json:"inbound"`
        
        Outbound []Mediatype `json:"outbound"`
        
        *Alias
    }{
        

        
        Inbound: []Mediatype{{}},
        

        

        
        Outbound: []Mediatype{{}},
        

        
        Alias: (*Alias)(u),
    })
}

