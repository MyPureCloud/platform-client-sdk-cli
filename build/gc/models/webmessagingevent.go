package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingeventDud struct { 
    


    


    

}

// Webmessagingevent - Message event element.  Examples include: system messages, typing indicators, cobrowse offerings.
type Webmessagingevent struct { 
    // EventType - Type of this event element
    EventType string `json:"eventType"`


    // CoBrowse - Cobrowse event.
    CoBrowse Webmessagingeventcobrowse `json:"coBrowse"`


    // Presence - Presence event.
    Presence Webmessagingeventpresence `json:"presence"`

}

// String returns a JSON representation of the model
func (o *Webmessagingevent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingevent) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingevent

    if WebmessagingeventMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingeventMarshalled = true

    return json.Marshal(&struct { 
        EventType string `json:"eventType"`
        
        CoBrowse Webmessagingeventcobrowse `json:"coBrowse"`
        
        Presence Webmessagingeventpresence `json:"presence"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

