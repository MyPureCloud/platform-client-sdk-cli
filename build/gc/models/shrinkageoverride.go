package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShrinkageoverrideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShrinkageoverrideDud struct { 
    


    

}

// Shrinkageoverride
type Shrinkageoverride struct { 
    // IntervalIndex - Index of shrinkage override interval. Starting index is 0 and indexes are based on 15 minute intervals for a 7 day week
    IntervalIndex int `json:"intervalIndex"`


    // ShrinkagePercent - Shrinkage override percent. Setting a null value will reset the interval to the default
    ShrinkagePercent float64 `json:"shrinkagePercent"`

}

// String returns a JSON representation of the model
func (o *Shrinkageoverride) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shrinkageoverride) MarshalJSON() ([]byte, error) {
    type Alias Shrinkageoverride

    if ShrinkageoverrideMarshalled {
        return []byte("{}"), nil
    }
    ShrinkageoverrideMarshalled = true

    return json.Marshal(&struct { 
        IntervalIndex int `json:"intervalIndex"`
        
        ShrinkagePercent float64 `json:"shrinkagePercent"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

