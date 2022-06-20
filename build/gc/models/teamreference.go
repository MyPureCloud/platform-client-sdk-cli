package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Teamreference
type Teamreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Teamreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamreference) MarshalJSON() ([]byte, error) {
    type Alias Teamreference

    if TeamreferenceMarshalled {
        return []byte("{}"), nil
    }
    TeamreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

