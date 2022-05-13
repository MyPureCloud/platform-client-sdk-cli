package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotificationsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotificationsresponseDud struct { 
    

}

// Notificationsresponse
type Notificationsresponse struct { 
    // Entities
    Entities []Wfmusernotification `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Notificationsresponse) String() string {
     o.Entities = []Wfmusernotification{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notificationsresponse) MarshalJSON() ([]byte, error) {
    type Alias Notificationsresponse

    if NotificationsresponseMarshalled {
        return []byte("{}"), nil
    }
    NotificationsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Wfmusernotification `json:"entities"`
        *Alias
    }{

        
        Entities: []Wfmusernotification{{}},
        

        Alias: (*Alias)(u),
    })
}

