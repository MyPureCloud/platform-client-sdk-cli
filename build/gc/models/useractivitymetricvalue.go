package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivitymetricvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivitymetricvalueDud struct { 
    


    

}

// Useractivitymetricvalue
type Useractivitymetricvalue struct { 
    // Metric - metric
    Metric string `json:"metric"`


    // Count - metric count
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Useractivitymetricvalue) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivitymetricvalue) MarshalJSON() ([]byte, error) {
    type Alias Useractivitymetricvalue

    if UseractivitymetricvalueMarshalled {
        return []byte("{}"), nil
    }
    UseractivitymetricvalueMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

