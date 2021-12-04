package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FailedrecordingentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FailedrecordingentitylistingDud struct { 
    

}

// Failedrecordingentitylisting
type Failedrecordingentitylisting struct { 
    // Entities
    Entities []Recordingjobfailedrecording `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Failedrecordingentitylisting) String() string {
    
    
     o.Entities = []Recordingjobfailedrecording{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Failedrecordingentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Failedrecordingentitylisting

    if FailedrecordingentitylistingMarshalled {
        return []byte("{}"), nil
    }
    FailedrecordingentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Recordingjobfailedrecording `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Recordingjobfailedrecording{{}},
        

        
        Alias: (*Alias)(u),
    })
}

