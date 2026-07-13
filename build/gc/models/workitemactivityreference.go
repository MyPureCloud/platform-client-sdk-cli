package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemactivityreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemactivityreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Workitemactivityreference
type Workitemactivityreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Workitemactivityreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemactivityreference) MarshalJSON() ([]byte, error) {
    type Alias Workitemactivityreference

    if WorkitemactivityreferenceMarshalled {
        return []byte("{}"), nil
    }
    WorkitemactivityreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

