package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuimporttimeofflimitvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuimporttimeofflimitvalueDud struct { 
    


    

}

// Buimporttimeofflimitvalue
type Buimporttimeofflimitvalue struct { 
    // ImportDateTime - The time-off limit interval UTC date time in ISO8601.
    ImportDateTime time.Time `json:"importDateTime"`


    // ImportMinutes - The limit value in minutes specified for a given date and time interval
    ImportMinutes int `json:"importMinutes"`

}

// String returns a JSON representation of the model
func (o *Buimporttimeofflimitvalue) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buimporttimeofflimitvalue) MarshalJSON() ([]byte, error) {
    type Alias Buimporttimeofflimitvalue

    if BuimporttimeofflimitvalueMarshalled {
        return []byte("{}"), nil
    }
    BuimporttimeofflimitvalueMarshalled = true

    return json.Marshal(&struct {
        
        ImportDateTime time.Time `json:"importDateTime"`
        
        ImportMinutes int `json:"importMinutes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

