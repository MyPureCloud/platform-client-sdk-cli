package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivitycodesreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivitycodesreferenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Activitycodesreference
type Activitycodesreference struct { 
    


    // Ids - The IDs of activity codes
    Ids []string `json:"ids"`


    

}

// String returns a JSON representation of the model
func (o *Activitycodesreference) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activitycodesreference) MarshalJSON() ([]byte, error) {
    type Alias Activitycodesreference

    if ActivitycodesreferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivitycodesreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        


        
        Ids: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

