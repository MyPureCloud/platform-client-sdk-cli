package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeofflimitreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeofflimitreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Timeofflimitreference
type Timeofflimitreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Timeofflimitreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeofflimitreference) MarshalJSON() ([]byte, error) {
    type Alias Timeofflimitreference

    if TimeofflimitreferenceMarshalled {
        return []byte("{}"), nil
    }
    TimeofflimitreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

