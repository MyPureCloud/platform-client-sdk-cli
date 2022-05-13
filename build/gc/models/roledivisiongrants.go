package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoledivisiongrantsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoledivisiongrantsDud struct { 
    

}

// Roledivisiongrants
type Roledivisiongrants struct { 
    // Grants - A list containing pairs of role and division IDs
    Grants []Roledivisionpair `json:"grants"`

}

// String returns a JSON representation of the model
func (o *Roledivisiongrants) String() string {
     o.Grants = []Roledivisionpair{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Roledivisiongrants) MarshalJSON() ([]byte, error) {
    type Alias Roledivisiongrants

    if RoledivisiongrantsMarshalled {
        return []byte("{}"), nil
    }
    RoledivisiongrantsMarshalled = true

    return json.Marshal(&struct {
        
        Grants []Roledivisionpair `json:"grants"`
        *Alias
    }{

        
        Grants: []Roledivisionpair{{}},
        

        Alias: (*Alias)(u),
    })
}

