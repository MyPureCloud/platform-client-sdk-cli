package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesdependentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesdependentDud struct { 
    Rule Usersrulesdependentrule `json:"rule"`


    TypeId string `json:"typeId"`


    VarType string `json:"type"`


    CreatedDate time.Time `json:"createdDate"`


    CreatedBy Userreference `json:"createdBy"`


    LastRun Usersruleslastrunmetadata `json:"lastRun"`


    RecentRunCount int `json:"recentRunCount"`


    SelfUri string `json:"selfUri"`

}

// Usersrulesdependent - Users rule dependent response
type Usersrulesdependent struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Usersrulesdependent) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesdependent) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesdependent

    if UsersrulesdependentMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesdependentMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

