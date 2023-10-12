package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventsessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventsessionDud struct { 
    


    


    

}

// Eventsession
type Eventsession struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // VarType - Session types indicate the type or category of sessions (e.g. web, app).
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Eventsession) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventsession) MarshalJSON() ([]byte, error) {
    type Alias Eventsession

    if EventsessionMarshalled {
        return []byte("{}"), nil
    }
    EventsessionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

