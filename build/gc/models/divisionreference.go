package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DivisionreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DivisionreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Divisionreference
type Divisionreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Divisionreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Divisionreference) MarshalJSON() ([]byte, error) {
    type Alias Divisionreference

    if DivisionreferenceMarshalled {
        return []byte("{}"), nil
    }
    DivisionreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

