package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StepreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StepreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Stepreference
type Stepreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Stepreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stepreference) MarshalJSON() ([]byte, error) {
    type Alias Stepreference

    if StepreferenceMarshalled {
        return []byte("{}"), nil
    }
    StepreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

