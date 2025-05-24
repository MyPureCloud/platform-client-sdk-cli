package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarysettingentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarysettingentityDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Summarysettingentity
type Summarysettingentity struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Summarysettingentity) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarysettingentity) MarshalJSON() ([]byte, error) {
    type Alias Summarysettingentity

    if SummarysettingentityMarshalled {
        return []byte("{}"), nil
    }
    SummarysettingentityMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

