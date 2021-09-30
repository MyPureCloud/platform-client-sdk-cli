package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParticipantattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParticipantattributesDud struct { 
    

}

// Participantattributes
type Participantattributes struct { 
    // Attributes - The map of attribute keys to values.
    Attributes map[string]string `json:"attributes"`

}

// String returns a JSON representation of the model
func (o *Participantattributes) String() string {
    
    
     o.Attributes = map[string]string{"": ""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Participantattributes) MarshalJSON() ([]byte, error) {
    type Alias Participantattributes

    if ParticipantattributesMarshalled {
        return []byte("{}"), nil
    }
    ParticipantattributesMarshalled = true

    return json.Marshal(&struct { 
        Attributes map[string]string `json:"attributes"`
        
        *Alias
    }{
        

        
        Attributes: map[string]string{"": ""},
        

        
        Alias: (*Alias)(u),
    })
}

