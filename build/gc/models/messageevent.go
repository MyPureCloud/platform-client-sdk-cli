package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessageeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessageeventDud struct { 
    


    


    


    


    


    

}

// Messageevent - Message event element.  Examples include: system messages, typing indicators, cobrowse offerings.
type Messageevent struct { 
    // EventType - Type of this event element
    EventType string `json:"eventType"`


    // CoBrowse - CoBrowse event.
    CoBrowse Eventcobrowse `json:"coBrowse"`


    // Typing - Typing event.
    Typing Eventtyping `json:"typing"`


    // Presence - Presence event.
    Presence Eventpresence `json:"presence"`


    // Video - Video event.
    Video Eventvideo `json:"video"`


    // Reactions - A list of reactions to a message.
    Reactions []Contentreaction `json:"reactions"`

}

// String returns a JSON representation of the model
func (o *Messageevent) String() string {
    
    
    
    
    
     o.Reactions = []Contentreaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messageevent) MarshalJSON() ([]byte, error) {
    type Alias Messageevent

    if MessageeventMarshalled {
        return []byte("{}"), nil
    }
    MessageeventMarshalled = true

    return json.Marshal(&struct {
        
        EventType string `json:"eventType"`
        
        CoBrowse Eventcobrowse `json:"coBrowse"`
        
        Typing Eventtyping `json:"typing"`
        
        Presence Eventpresence `json:"presence"`
        
        Video Eventvideo `json:"video"`
        
        Reactions []Contentreaction `json:"reactions"`
        *Alias
    }{

        


        


        


        


        


        
        Reactions: []Contentreaction{{}},
        

        Alias: (*Alias)(u),
    })
}

