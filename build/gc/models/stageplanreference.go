package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StageplanreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StageplanreferenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Stageplanreference
type Stageplanreference struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Stageplanreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stageplanreference) MarshalJSON() ([]byte, error) {
    type Alias Stageplanreference

    if StageplanreferenceMarshalled {
        return []byte("{}"), nil
    }
    StageplanreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

