package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListwrapperbidgroupworkplanrotationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListwrapperbidgroupworkplanrotationrequestDud struct { 
    

}

// Listwrapperbidgroupworkplanrotationrequest
type Listwrapperbidgroupworkplanrotationrequest struct { 
    // Values
    Values []Bidgroupworkplanrotationrequest `json:"values"`

}

// String returns a JSON representation of the model
func (o *Listwrapperbidgroupworkplanrotationrequest) String() string {
     o.Values = []Bidgroupworkplanrotationrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listwrapperbidgroupworkplanrotationrequest) MarshalJSON() ([]byte, error) {
    type Alias Listwrapperbidgroupworkplanrotationrequest

    if ListwrapperbidgroupworkplanrotationrequestMarshalled {
        return []byte("{}"), nil
    }
    ListwrapperbidgroupworkplanrotationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Values []Bidgroupworkplanrotationrequest `json:"values"`
        *Alias
    }{

        
        Values: []Bidgroupworkplanrotationrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

