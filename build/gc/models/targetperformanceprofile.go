package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TargetperformanceprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TargetperformanceprofileDud struct { 
    

}

// Targetperformanceprofile
type Targetperformanceprofile struct { 
    // TargetPerformanceProfileId - The target destination performanceProfileId for the linked metric.
    TargetPerformanceProfileId string `json:"targetPerformanceProfileId"`

}

// String returns a JSON representation of the model
func (o *Targetperformanceprofile) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Targetperformanceprofile) MarshalJSON() ([]byte, error) {
    type Alias Targetperformanceprofile

    if TargetperformanceprofileMarshalled {
        return []byte("{}"), nil
    }
    TargetperformanceprofileMarshalled = true

    return json.Marshal(&struct {
        
        TargetPerformanceProfileId string `json:"targetPerformanceProfileId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

