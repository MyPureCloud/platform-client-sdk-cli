package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemqueuereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemqueuereferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Workitemqueuereference
type Workitemqueuereference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Workitemqueuereference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemqueuereference) MarshalJSON() ([]byte, error) {
    type Alias Workitemqueuereference

    if WorkitemqueuereferenceMarshalled {
        return []byte("{}"), nil
    }
    WorkitemqueuereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

