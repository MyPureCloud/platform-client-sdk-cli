package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RangeDud struct { 
    


    


    

}

// Range
type Range struct { 
    // VarType - Range type (NoEnd: without an end date. EndDate: with an end date. Numbered: with a specific number of occurrences)
    VarType string `json:"type"`


    // End - The end date time of the last occurrence of the range as an ISO-8601 string in UTC time, e.g: 2023-12-21T16:30:25.000Z, Required to set for EndDate range type.
    End string `json:"end"`


    // NumberOfOccurrences - The number of times the schedule will be repeated, e.g: 2. Required to set for Numbered range type.
    NumberOfOccurrences int `json:"numberOfOccurrences"`

}

// String returns a JSON representation of the model
func (o *Range) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Range) MarshalJSON() ([]byte, error) {
    type Alias Range

    if RangeMarshalled {
        return []byte("{}"), nil
    }
    RangeMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        End string `json:"end"`
        
        NumberOfOccurrences int `json:"numberOfOccurrences"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

