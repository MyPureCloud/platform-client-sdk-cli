package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsertimeoffintegrationstatusresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsertimeoffintegrationstatusresponselistingDud struct { 
    

}

// Usertimeoffintegrationstatusresponselisting
type Usertimeoffintegrationstatusresponselisting struct { 
    // Entities
    Entities []Usertimeoffintegrationstatusresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Usertimeoffintegrationstatusresponselisting) String() string {
     o.Entities = []Usertimeoffintegrationstatusresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usertimeoffintegrationstatusresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Usertimeoffintegrationstatusresponselisting

    if UsertimeoffintegrationstatusresponselistingMarshalled {
        return []byte("{}"), nil
    }
    UsertimeoffintegrationstatusresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Usertimeoffintegrationstatusresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Usertimeoffintegrationstatusresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

