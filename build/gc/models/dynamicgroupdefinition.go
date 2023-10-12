package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamicgroupdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamicgroupdefinitionDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Dynamicgroupdefinition
type Dynamicgroupdefinition struct { 
    // Group - The group the dynamic group definition belongs to
    Group Addressableentityref `json:"group"`


    // DateModified - Last modified date/time of the dynamic group definition. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Query - Properties used for building a dynamic groups query
    Query Dynamicgroupquery `json:"query"`


    

}

// String returns a JSON representation of the model
func (o *Dynamicgroupdefinition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamicgroupdefinition) MarshalJSON() ([]byte, error) {
    type Alias Dynamicgroupdefinition

    if DynamicgroupdefinitionMarshalled {
        return []byte("{}"), nil
    }
    DynamicgroupdefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Group Addressableentityref `json:"group"`
        
        DateModified time.Time `json:"dateModified"`
        
        Query Dynamicgroupquery `json:"query"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

