package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidgroupreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidgroupreferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Schedulebidgroupreference
type Schedulebidgroupreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Schedulebidgroupreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidgroupreference) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidgroupreference

    if SchedulebidgroupreferenceMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidgroupreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

