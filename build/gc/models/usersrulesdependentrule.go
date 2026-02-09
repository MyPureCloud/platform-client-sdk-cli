package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesdependentruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesdependentruleDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Usersrulesdependentrule - Users rule response
type Usersrulesdependentrule struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Usersrulesdependentrule) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesdependentrule) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesdependentrule

    if UsersrulesdependentruleMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesdependentruleMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

