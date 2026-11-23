package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitieslimitvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitieslimitvalueDud struct { 
    IntegerValue int `json:"integerValue"`


    StringValue string `json:"stringValue"`

}

// Capabilitieslimitvalue
type Capabilitieslimitvalue struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Capabilitieslimitvalue) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitieslimitvalue) MarshalJSON() ([]byte, error) {
    type Alias Capabilitieslimitvalue

    if CapabilitieslimitvalueMarshalled {
        return []byte("{}"), nil
    }
    CapabilitieslimitvalueMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

