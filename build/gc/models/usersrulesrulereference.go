package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesrulereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesrulereferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Usersrulesrulereference
type Usersrulesrulereference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Usersrulesrulereference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesrulereference) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesrulereference

    if UsersrulesrulereferenceMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesrulereferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

