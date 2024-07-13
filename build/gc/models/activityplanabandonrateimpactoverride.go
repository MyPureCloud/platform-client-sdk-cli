package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanabandonrateimpactoverrideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanabandonrateimpactoverrideDud struct { 
    

}

// Activityplanabandonrateimpactoverride
type Activityplanabandonrateimpactoverride struct { 
    // IncreaseByPercent - Allowed abandon rate increase percent, from 0.0 to 100.0
    IncreaseByPercent float64 `json:"increaseByPercent"`

}

// String returns a JSON representation of the model
func (o *Activityplanabandonrateimpactoverride) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanabandonrateimpactoverride) MarshalJSON() ([]byte, error) {
    type Alias Activityplanabandonrateimpactoverride

    if ActivityplanabandonrateimpactoverrideMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanabandonrateimpactoverrideMarshalled = true

    return json.Marshal(&struct {
        
        IncreaseByPercent float64 `json:"increaseByPercent"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

