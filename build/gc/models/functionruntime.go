package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FunctionruntimeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FunctionruntimeDud struct { 
    Name string `json:"name"`


    Description string `json:"description"`


    Status string `json:"status"`


    DateEndOfLife time.Time `json:"dateEndOfLife"`

}

// Functionruntime - Action function runtime.
type Functionruntime struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Functionruntime) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Functionruntime) MarshalJSON() ([]byte, error) {
    type Alias Functionruntime

    if FunctionruntimeMarshalled {
        return []byte("{}"), nil
    }
    FunctionruntimeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

