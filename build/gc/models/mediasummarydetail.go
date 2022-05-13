package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediasummarydetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediasummarydetailDud struct { 
    


    

}

// Mediasummarydetail
type Mediasummarydetail struct { 
    // Active
    Active int `json:"active"`


    // Acw
    Acw int `json:"acw"`

}

// String returns a JSON representation of the model
func (o *Mediasummarydetail) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediasummarydetail) MarshalJSON() ([]byte, error) {
    type Alias Mediasummarydetail

    if MediasummarydetailMarshalled {
        return []byte("{}"), nil
    }
    MediasummarydetailMarshalled = true

    return json.Marshal(&struct {
        
        Active int `json:"active"`
        
        Acw int `json:"acw"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

