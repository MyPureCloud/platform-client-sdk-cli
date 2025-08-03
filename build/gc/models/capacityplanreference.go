package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplanreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplanreferenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Capacityplanreference
type Capacityplanreference struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Capacityplanreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplanreference) MarshalJSON() ([]byte, error) {
    type Alias Capacityplanreference

    if CapacityplanreferenceMarshalled {
        return []byte("{}"), nil
    }
    CapacityplanreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

