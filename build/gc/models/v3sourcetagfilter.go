package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcetagfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcetagfilterDud struct { 
    


    


    

}

// V3sourcetagfilter
type V3sourcetagfilter struct { 
    // AllOf - Tags that must all be present (AND).
    AllOf []string `json:"allOf"`


    // AnyOfGroups - OR groups ANDed together; within each group at least one tag must match.
    AnyOfGroups [][]string `json:"anyOfGroups"`


    // NoneOf - Tags that must not be present on matching chunks.
    NoneOf []string `json:"noneOf"`

}

// String returns a JSON representation of the model
func (o *V3sourcetagfilter) String() string {
     o.AllOf = []string{""} 
     o.AnyOfGroups = [][]string{{}} 
     o.NoneOf = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcetagfilter) MarshalJSON() ([]byte, error) {
    type Alias V3sourcetagfilter

    if V3sourcetagfilterMarshalled {
        return []byte("{}"), nil
    }
    V3sourcetagfilterMarshalled = true

    return json.Marshal(&struct {
        
        AllOf []string `json:"allOf"`
        
        AnyOfGroups [][]string `json:"anyOfGroups"`
        
        NoneOf []string `json:"noneOf"`
        *Alias
    }{

        
        AllOf: []string{""},
        


        
        AnyOfGroups: [][]string{{}},
        


        
        NoneOf: []string{""},
        

        Alias: (*Alias)(u),
    })
}

