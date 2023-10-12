package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamicgroupquerypreviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamicgroupquerypreviewDud struct { 
    

}

// Dynamicgroupquerypreview
type Dynamicgroupquerypreview struct { 
    // UserCount - Number of Users that match the DynamicGroupQuery
    UserCount int `json:"userCount"`

}

// String returns a JSON representation of the model
func (o *Dynamicgroupquerypreview) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamicgroupquerypreview) MarshalJSON() ([]byte, error) {
    type Alias Dynamicgroupquerypreview

    if DynamicgroupquerypreviewMarshalled {
        return []byte("{}"), nil
    }
    DynamicgroupquerypreviewMarshalled = true

    return json.Marshal(&struct {
        
        UserCount int `json:"userCount"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

