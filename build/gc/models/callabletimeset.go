package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallabletimesetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallabletimesetDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    SelfUri string `json:"selfUri"`

}

// Callabletimeset
type Callabletimeset struct { 
    


    // Name - The name of the CallableTimeSet.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // CallableTimes - The list of CallableTimes for which it is acceptable to place outbound calls.
    CallableTimes []Callabletime `json:"callableTimes"`


    

}

// String returns a JSON representation of the model
func (o *Callabletimeset) String() string {
    
    
     o.CallableTimes = []Callabletime{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callabletimeset) MarshalJSON() ([]byte, error) {
    type Alias Callabletimeset

    if CallabletimesetMarshalled {
        return []byte("{}"), nil
    }
    CallabletimesetMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        CallableTimes []Callabletime `json:"callableTimes"`
        *Alias
    }{

        


        


        


        


        


        
        CallableTimes: []Callabletime{{}},
        


        

        Alias: (*Alias)(u),
    })
}

