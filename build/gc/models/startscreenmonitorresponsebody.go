package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StartscreenmonitorresponsebodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StartscreenmonitorresponsebodyDud struct { 
    

}

// Startscreenmonitorresponsebody
type Startscreenmonitorresponsebody struct { 
    // ScreenMonitoringId
    ScreenMonitoringId string `json:"screenMonitoringId"`

}

// String returns a JSON representation of the model
func (o *Startscreenmonitorresponsebody) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Startscreenmonitorresponsebody) MarshalJSON() ([]byte, error) {
    type Alias Startscreenmonitorresponsebody

    if StartscreenmonitorresponsebodyMarshalled {
        return []byte("{}"), nil
    }
    StartscreenmonitorresponsebodyMarshalled = true

    return json.Marshal(&struct {
        
        ScreenMonitoringId string `json:"screenMonitoringId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

