package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulereferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Usersrulereference
type Usersrulereference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Usersrulereference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulereference) MarshalJSON() ([]byte, error) {
    type Alias Usersrulereference

    if UsersrulereferenceMarshalled {
        return []byte("{}"), nil
    }
    UsersrulereferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

