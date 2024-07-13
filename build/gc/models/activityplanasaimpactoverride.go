package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanasaimpactoverrideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanasaimpactoverrideDud struct { 
    

}

// Activityplanasaimpactoverride
type Activityplanasaimpactoverride struct { 
    // IncreaseByPercent - Allowed average speed of answer increase percent, from 0.0 to 100.0
    IncreaseByPercent float64 `json:"increaseByPercent"`

}

// String returns a JSON representation of the model
func (o *Activityplanasaimpactoverride) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanasaimpactoverride) MarshalJSON() ([]byte, error) {
    type Alias Activityplanasaimpactoverride

    if ActivityplanasaimpactoverrideMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanasaimpactoverrideMarshalled = true

    return json.Marshal(&struct {
        
        IncreaseByPercent float64 `json:"increaseByPercent"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

