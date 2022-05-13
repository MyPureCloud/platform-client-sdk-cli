package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserbestpointsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserbestpointsDud struct { 
    User Userreference `json:"user"`


    BestPoints []Userbestpointsitem `json:"bestPoints"`

}

// Userbestpoints
type Userbestpoints struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Userbestpoints) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userbestpoints) MarshalJSON() ([]byte, error) {
    type Alias Userbestpoints

    if UserbestpointsMarshalled {
        return []byte("{}"), nil
    }
    UserbestpointsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

