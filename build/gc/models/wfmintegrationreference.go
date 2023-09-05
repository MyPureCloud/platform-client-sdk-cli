package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmintegrationreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmintegrationreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Wfmintegrationreference
type Wfmintegrationreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Wfmintegrationreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmintegrationreference) MarshalJSON() ([]byte, error) {
    type Alias Wfmintegrationreference

    if WfmintegrationreferenceMarshalled {
        return []byte("{}"), nil
    }
    WfmintegrationreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

