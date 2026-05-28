package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaleventchangesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaleventchangesresponseDud struct { 
    Changes []Externaleventchange `json:"changes"`

}

// Externaleventchangesresponse - Response for getting changes in external event definitions
type Externaleventchangesresponse struct { 
    

}

// String returns a JSON representation of the model
func (o *Externaleventchangesresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externaleventchangesresponse) MarshalJSON() ([]byte, error) {
    type Alias Externaleventchangesresponse

    if ExternaleventchangesresponseMarshalled {
        return []byte("{}"), nil
    }
    ExternaleventchangesresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

