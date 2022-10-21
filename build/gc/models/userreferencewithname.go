package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserreferencewithnameMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserreferencewithnameDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Userreferencewithname
type Userreferencewithname struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Userreferencewithname) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userreferencewithname) MarshalJSON() ([]byte, error) {
    type Alias Userreferencewithname

    if UserreferencewithnameMarshalled {
        return []byte("{}"), nil
    }
    UserreferencewithnameMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

