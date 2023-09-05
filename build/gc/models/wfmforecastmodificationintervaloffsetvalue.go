package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmforecastmodificationintervaloffsetvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmforecastmodificationintervaloffsetvalueDud struct { 
    


    

}

// Wfmforecastmodificationintervaloffsetvalue
type Wfmforecastmodificationintervaloffsetvalue struct { 
    // IntervalIndex - The number of intervals past referenceStartDate to which to apply this modification
    IntervalIndex int `json:"intervalIndex"`


    // Value - The value to set for the given interval
    Value float64 `json:"value"`

}

// String returns a JSON representation of the model
func (o *Wfmforecastmodificationintervaloffsetvalue) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmforecastmodificationintervaloffsetvalue) MarshalJSON() ([]byte, error) {
    type Alias Wfmforecastmodificationintervaloffsetvalue

    if WfmforecastmodificationintervaloffsetvalueMarshalled {
        return []byte("{}"), nil
    }
    WfmforecastmodificationintervaloffsetvalueMarshalled = true

    return json.Marshal(&struct {
        
        IntervalIndex int `json:"intervalIndex"`
        
        Value float64 `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

