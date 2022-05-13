package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatenotificationsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatenotificationsresponseDud struct { 
    

}

// Updatenotificationsresponse
type Updatenotificationsresponse struct { 
    // Entities
    Entities []Updatenotificationresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Updatenotificationsresponse) String() string {
     o.Entities = []Updatenotificationresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatenotificationsresponse) MarshalJSON() ([]byte, error) {
    type Alias Updatenotificationsresponse

    if UpdatenotificationsresponseMarshalled {
        return []byte("{}"), nil
    }
    UpdatenotificationsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Updatenotificationresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Updatenotificationresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

