package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemaccessattributesupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemaccessattributesupdateDud struct { 
    

}

// Workitemaccessattributesupdate
type Workitemaccessattributesupdate struct { 
    // AccessAttributes - The attributes that define which users may access a workitem.
    AccessAttributes []string `json:"accessAttributes"`

}

// String returns a JSON representation of the model
func (o *Workitemaccessattributesupdate) String() string {
     o.AccessAttributes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemaccessattributesupdate) MarshalJSON() ([]byte, error) {
    type Alias Workitemaccessattributesupdate

    if WorkitemaccessattributesupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemaccessattributesupdateMarshalled = true

    return json.Marshal(&struct {
        
        AccessAttributes []string `json:"accessAttributes"`
        *Alias
    }{

        
        AccessAttributes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

