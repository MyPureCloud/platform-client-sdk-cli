package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuideversionrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuideversionrefDud struct { 
    Version string `json:"version"`


    SelfUri string `json:"selfUri"`

}

// Guideversionref
type Guideversionref struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Guideversionref) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guideversionref) MarshalJSON() ([]byte, error) {
    type Alias Guideversionref

    if GuideversionrefMarshalled {
        return []byte("{}"), nil
    }
    GuideversionrefMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

