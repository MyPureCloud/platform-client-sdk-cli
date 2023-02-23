package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomerefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomerefDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Outcomeref
type Outcomeref struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Outcomeref) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeref) MarshalJSON() ([]byte, error) {
    type Alias Outcomeref

    if OutcomerefMarshalled {
        return []byte("{}"), nil
    }
    OutcomerefMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

