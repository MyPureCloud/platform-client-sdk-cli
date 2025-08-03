package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DefaultanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DefaultanswerDud struct { 
    


    NotApplicable bool `json:"notApplicable"`

}

// Defaultanswer
type Defaultanswer struct { 
    // Id - Selected default answer id for the question.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Defaultanswer) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Defaultanswer) MarshalJSON() ([]byte, error) {
    type Alias Defaultanswer

    if DefaultanswerMarshalled {
        return []byte("{}"), nil
    }
    DefaultanswerMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

