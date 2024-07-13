package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivitycodereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivitycodereferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Activitycodereference
type Activitycodereference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Activitycodereference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activitycodereference) MarshalJSON() ([]byte, error) {
    type Alias Activitycodereference

    if ActivitycodereferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivitycodereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

