package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediareactionseventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediareactionseventDud struct { 
    

}

// Opensocialmediareactionsevent - Social Message event element.
type Opensocialmediareactionsevent struct { 
    // Reactions - List of reactions for this event.
    Reactions []Contentreaction `json:"reactions"`

}

// String returns a JSON representation of the model
func (o *Opensocialmediareactionsevent) String() string {
     o.Reactions = []Contentreaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediareactionsevent) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediareactionsevent

    if OpensocialmediareactionseventMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediareactionseventMarshalled = true

    return json.Marshal(&struct {
        
        Reactions []Contentreaction `json:"reactions"`
        *Alias
    }{

        
        Reactions: []Contentreaction{{}},
        

        Alias: (*Alias)(u),
    })
}

