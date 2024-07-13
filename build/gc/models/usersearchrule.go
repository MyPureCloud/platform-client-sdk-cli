package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersearchruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersearchruleDud struct { 
    

}

// Usersearchrule
type Usersearchrule struct { 
    // Parts - The parts of this rule; the results of these are ANDed together.
    Parts []Usersearchrulepart `json:"parts"`

}

// String returns a JSON representation of the model
func (o *Usersearchrule) String() string {
     o.Parts = []Usersearchrulepart{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersearchrule) MarshalJSON() ([]byte, error) {
    type Alias Usersearchrule

    if UsersearchruleMarshalled {
        return []byte("{}"), nil
    }
    UsersearchruleMarshalled = true

    return json.Marshal(&struct {
        
        Parts []Usersearchrulepart `json:"parts"`
        *Alias
    }{

        
        Parts: []Usersearchrulepart{{}},
        

        Alias: (*Alias)(u),
    })
}

