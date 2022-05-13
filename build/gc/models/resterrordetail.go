package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResterrordetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResterrordetailDud struct { 
    VarError string `json:"error"`


    Details string `json:"details"`

}

// Resterrordetail
type Resterrordetail struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Resterrordetail) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resterrordetail) MarshalJSON() ([]byte, error) {
    type Alias Resterrordetail

    if ResterrordetailMarshalled {
        return []byte("{}"), nil
    }
    ResterrordetailMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

