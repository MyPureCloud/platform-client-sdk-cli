package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeofflimitlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeofflimitlistingDud struct { 
    

}

// Timeofflimitlisting - The list of time off limit objects
type Timeofflimitlisting struct { 
    // Entities
    Entities []Timeofflimit `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Timeofflimitlisting) String() string {
    
    
     o.Entities = []Timeofflimit{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeofflimitlisting) MarshalJSON() ([]byte, error) {
    type Alias Timeofflimitlisting

    if TimeofflimitlistingMarshalled {
        return []byte("{}"), nil
    }
    TimeofflimitlistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Timeofflimit `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Timeofflimit{{}},
        

        
        Alias: (*Alias)(u),
    })
}

