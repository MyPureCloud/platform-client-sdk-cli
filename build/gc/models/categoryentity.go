package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CategoryentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CategoryentityDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Categoryentity
type Categoryentity struct { 
    // Id - The Id of the category.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Categoryentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Categoryentity) MarshalJSON() ([]byte, error) {
    type Alias Categoryentity

    if CategoryentityMarshalled {
        return []byte("{}"), nil
    }
    CategoryentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

