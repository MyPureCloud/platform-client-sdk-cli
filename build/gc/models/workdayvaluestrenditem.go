package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkdayvaluestrenditemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkdayvaluestrenditemDud struct { 
    DateWorkday time.Time `json:"dateWorkday"`


    Value float64 `json:"value"`

}

// Workdayvaluestrenditem
type Workdayvaluestrenditem struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Workdayvaluestrenditem) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workdayvaluestrenditem) MarshalJSON() ([]byte, error) {
    type Alias Workdayvaluestrenditem

    if WorkdayvaluestrenditemMarshalled {
        return []byte("{}"), nil
    }
    WorkdayvaluestrenditemMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

