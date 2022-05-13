package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectedqueueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectedqueueDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Connectedqueue
type Connectedqueue struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Connectedqueue) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectedqueue) MarshalJSON() ([]byte, error) {
    type Alias Connectedqueue

    if ConnectedqueueMarshalled {
        return []byte("{}"), nil
    }
    ConnectedqueueMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

