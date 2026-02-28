package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcerefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcerefDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// V3sourceref
type V3sourceref struct { 
    // Id - Source Id.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *V3sourceref) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourceref) MarshalJSON() ([]byte, error) {
    type Alias V3sourceref

    if V3sourcerefMarshalled {
        return []byte("{}"), nil
    }
    V3sourcerefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

