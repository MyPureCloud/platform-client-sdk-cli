package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanservicelevelimpactoverrideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanservicelevelimpactoverrideDud struct { 
    

}

// Activityplanservicelevelimpactoverride
type Activityplanservicelevelimpactoverride struct { 
    // DecreaseByPercent - Allowed service level decrease percent, from 0.0 to 100.0
    DecreaseByPercent float64 `json:"decreaseByPercent"`

}

// String returns a JSON representation of the model
func (o *Activityplanservicelevelimpactoverride) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanservicelevelimpactoverride) MarshalJSON() ([]byte, error) {
    type Alias Activityplanservicelevelimpactoverride

    if ActivityplanservicelevelimpactoverrideMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanservicelevelimpactoverrideMarshalled = true

    return json.Marshal(&struct {
        
        DecreaseByPercent float64 `json:"decreaseByPercent"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

