package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OverridedateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OverridedateDud struct { 
    


    

}

// Overridedate
type Overridedate struct { 
    // Date - The date in yyyy-MM-dd format, interpreted in the business unit’s time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    Date time.Time `json:"date"`


    // VarType - The type of override date
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Overridedate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Overridedate) MarshalJSON() ([]byte, error) {
    type Alias Overridedate

    if OverridedateMarshalled {
        return []byte("{}"), nil
    }
    OverridedateMarshalled = true

    return json.Marshal(&struct {
        
        Date time.Time `json:"date"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

