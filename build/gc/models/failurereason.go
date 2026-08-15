package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FailurereasonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FailurereasonDud struct { 
    


    

}

// Failurereason
type Failurereason struct { 
    // Code - System-defined failure code.
    Code string `json:"code"`


    // Message - Human-readable description of the failure
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Failurereason) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Failurereason) MarshalJSON() ([]byte, error) {
    type Alias Failurereason

    if FailurereasonMarshalled {
        return []byte("{}"), nil
    }
    FailurereasonMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

