package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnansweredgroupsuggesteddocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnansweredgroupsuggesteddocumentDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Unansweredgroupsuggesteddocument
type Unansweredgroupsuggesteddocument struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Unansweredgroupsuggesteddocument) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unansweredgroupsuggesteddocument) MarshalJSON() ([]byte, error) {
    type Alias Unansweredgroupsuggesteddocument

    if UnansweredgroupsuggesteddocumentMarshalled {
        return []byte("{}"), nil
    }
    UnansweredgroupsuggesteddocumentMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

