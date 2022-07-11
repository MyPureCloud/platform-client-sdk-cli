package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShiftstartvarianceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShiftstartvarianceDud struct { 
    


    

}

// Shiftstartvariance
type Shiftstartvariance struct { 
    // ApplicableDays - Days for which shift start variance is configured
    ApplicableDays []string `json:"applicableDays"`


    // MaxShiftStartVarianceMinutes - Maximum variance in minutes across shift starts
    MaxShiftStartVarianceMinutes int `json:"maxShiftStartVarianceMinutes"`

}

// String returns a JSON representation of the model
func (o *Shiftstartvariance) String() string {
     o.ApplicableDays = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shiftstartvariance) MarshalJSON() ([]byte, error) {
    type Alias Shiftstartvariance

    if ShiftstartvarianceMarshalled {
        return []byte("{}"), nil
    }
    ShiftstartvarianceMarshalled = true

    return json.Marshal(&struct {
        
        ApplicableDays []string `json:"applicableDays"`
        
        MaxShiftStartVarianceMinutes int `json:"maxShiftStartVarianceMinutes"`
        *Alias
    }{

        
        ApplicableDays: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

