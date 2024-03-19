package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeofflimitreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeofflimitreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Butimeofflimitreference
type Butimeofflimitreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Butimeofflimitreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeofflimitreference) MarshalJSON() ([]byte, error) {
    type Alias Butimeofflimitreference

    if ButimeofflimitreferenceMarshalled {
        return []byte("{}"), nil
    }
    ButimeofflimitreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

