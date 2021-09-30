package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatenotificationsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatenotificationsrequestDud struct { 
    

}

// Updatenotificationsrequest
type Updatenotificationsrequest struct { 
    // Entities - The notifications to update
    Entities []Wfmusernotification `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Updatenotificationsrequest) String() string {
    
    
     o.Entities = []Wfmusernotification{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatenotificationsrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatenotificationsrequest

    if UpdatenotificationsrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatenotificationsrequestMarshalled = true

    return json.Marshal(&struct { 
        Entities []Wfmusernotification `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Wfmusernotification{{}},
        

        
        Alias: (*Alias)(u),
    })
}

