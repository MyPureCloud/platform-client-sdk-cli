package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingjobentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingjobentitylistingDud struct { 
    

}

// Recordingjobentitylisting
type Recordingjobentitylisting struct { 
    // Entities
    Entities []Recordingjob `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Recordingjobentitylisting) String() string {
    
    
     o.Entities = []Recordingjob{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingjobentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Recordingjobentitylisting

    if RecordingjobentitylistingMarshalled {
        return []byte("{}"), nil
    }
    RecordingjobentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Recordingjob `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Recordingjob{{}},
        

        
        Alias: (*Alias)(u),
    })
}

