package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Activityplanreference
type Activityplanreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Activityplanreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanreference) MarshalJSON() ([]byte, error) {
    type Alias Activityplanreference

    if ActivityplanreferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

