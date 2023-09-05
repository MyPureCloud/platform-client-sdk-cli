package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffintegrationstatusresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffintegrationstatusresponselistingDud struct { 
    

}

// Timeoffintegrationstatusresponselisting
type Timeoffintegrationstatusresponselisting struct { 
    // Entities
    Entities []Timeoffintegrationstatusresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Timeoffintegrationstatusresponselisting) String() string {
     o.Entities = []Timeoffintegrationstatusresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffintegrationstatusresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Timeoffintegrationstatusresponselisting

    if TimeoffintegrationstatusresponselistingMarshalled {
        return []byte("{}"), nil
    }
    TimeoffintegrationstatusresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Timeoffintegrationstatusresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Timeoffintegrationstatusresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

