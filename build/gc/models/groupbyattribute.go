package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GroupbyattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GroupbyattributeDud struct { 
    Attribute string `json:"attribute"`


    Value string `json:"value"`

}

// Groupbyattribute
type Groupbyattribute struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Groupbyattribute) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Groupbyattribute) MarshalJSON() ([]byte, error) {
    type Alias Groupbyattribute

    if GroupbyattributeMarshalled {
        return []byte("{}"), nil
    }
    GroupbyattributeMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

