package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasereferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Casereference
type Casereference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Casereference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casereference) MarshalJSON() ([]byte, error) {
    type Alias Casereference

    if CasereferenceMarshalled {
        return []byte("{}"), nil
    }
    CasereferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

