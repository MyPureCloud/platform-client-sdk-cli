package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffinggroupreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffinggroupreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Staffinggroupreference
type Staffinggroupreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Staffinggroupreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffinggroupreference) MarshalJSON() ([]byte, error) {
    type Alias Staffinggroupreference

    if StaffinggroupreferenceMarshalled {
        return []byte("{}"), nil
    }
    StaffinggroupreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

