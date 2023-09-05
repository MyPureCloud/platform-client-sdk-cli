package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenmessageeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenmessageeventDud struct { 
    


    

}

// Openmessageevent - Message event element.
type Openmessageevent struct { 
    // EventType - Type of this event element
    EventType string `json:"eventType"`


    // Typing - Typing event.
    Typing Conversationeventtyping `json:"typing"`

}

// String returns a JSON representation of the model
func (o *Openmessageevent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openmessageevent) MarshalJSON() ([]byte, error) {
    type Alias Openmessageevent

    if OpenmessageeventMarshalled {
        return []byte("{}"), nil
    }
    OpenmessageeventMarshalled = true

    return json.Marshal(&struct {
        
        EventType string `json:"eventType"`
        
        Typing Conversationeventtyping `json:"typing"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

