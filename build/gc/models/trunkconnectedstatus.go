package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkconnectedstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkconnectedstatusDud struct { 
    


    

}

// Trunkconnectedstatus
type Trunkconnectedstatus struct { 
    // Connected
    Connected bool `json:"connected"`


    // ConnectedStateTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedStateTime time.Time `json:"connectedStateTime"`

}

// String returns a JSON representation of the model
func (o *Trunkconnectedstatus) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkconnectedstatus) MarshalJSON() ([]byte, error) {
    type Alias Trunkconnectedstatus

    if TrunkconnectedstatusMarshalled {
        return []byte("{}"), nil
    }
    TrunkconnectedstatusMarshalled = true

    return json.Marshal(&struct {
        
        Connected bool `json:"connected"`
        
        ConnectedStateTime time.Time `json:"connectedStateTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

