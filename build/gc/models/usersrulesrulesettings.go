package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesrulesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesrulesettingsDud struct { 
    AllowedContainers []string `json:"allowedContainers"`


    LockedCriteria []Usersruleslockedcriteriasettingscriteria `json:"lockedCriteria"`

}

// Usersrulesrulesettings - Users rule type-specific settings response
type Usersrulesrulesettings struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Usersrulesrulesettings) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesrulesettings) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesrulesettings

    if UsersrulesrulesettingsMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesrulesettingsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

