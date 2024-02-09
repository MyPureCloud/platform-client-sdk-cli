package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignedsegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignedsegmentDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Assignedsegment
type Assignedsegment struct { 
    // Id - The ID of the segment assigned.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Assignedsegment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignedsegment) MarshalJSON() ([]byte, error) {
    type Alias Assignedsegment

    if AssignedsegmentMarshalled {
        return []byte("{}"), nil
    }
    AssignedsegmentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

