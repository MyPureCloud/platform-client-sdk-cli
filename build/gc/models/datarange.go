package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DatarangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DatarangeDud struct { 
    DateMin time.Time `json:"dateMin"`


    DateMax time.Time `json:"dateMax"`

}

// Datarange
type Datarange struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Datarange) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Datarange) MarshalJSON() ([]byte, error) {
    type Alias Datarange

    if DatarangeMarshalled {
        return []byte("{}"), nil
    }
    DatarangeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

