package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanrotationreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanrotationreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Workplanrotationreference
type Workplanrotationreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Workplanrotationreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanrotationreference) MarshalJSON() ([]byte, error) {
    type Alias Workplanrotationreference

    if WorkplanrotationreferenceMarshalled {
        return []byte("{}"), nil
    }
    WorkplanrotationreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

