package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NumberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NumberDud struct { 
    


    

}

// Number
type Number struct { 
    // Start
    Start string `json:"start"`


    // End
    End string `json:"end"`

}

// String returns a JSON representation of the model
func (o *Number) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Number) MarshalJSON() ([]byte, error) {
    type Alias Number

    if NumberMarshalled {
        return []byte("{}"), nil
    }
    NumberMarshalled = true

    return json.Marshal(&struct {
        
        Start string `json:"start"`
        
        End string `json:"end"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

