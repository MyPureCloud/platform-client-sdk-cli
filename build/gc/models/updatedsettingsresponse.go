package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatedsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatedsettingsresponseDud struct { 
    

}

// Updatedsettingsresponse
type Updatedsettingsresponse struct { 
    // Message
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Updatedsettingsresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatedsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Updatedsettingsresponse

    if UpdatedsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    UpdatedsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

