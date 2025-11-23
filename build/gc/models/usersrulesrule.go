package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesruleDud struct { 
    Id string `json:"id"`


    


    Description string `json:"description"`


    VarType string `json:"type"`


    Criteria []Usersrulescriteria `json:"criteria"`


    CreatedDate time.Time `json:"createdDate"`


    CreatedBy Userreference `json:"createdBy"`


    ModifiedDate time.Time `json:"modifiedDate"`


    ModifiedBy Userreference `json:"modifiedBy"`


    LastRun Usersruleslastrunmetadata `json:"lastRun"`


    RecentRunCount int `json:"recentRunCount"`


    DependentCount int `json:"dependentCount"`


    SelfUri string `json:"selfUri"`

}

// Usersrulesrule - Users rule response
type Usersrulesrule struct { 
    


    // Name
    Name string `json:"name"`


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Usersrulesrule) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesrule) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesrule

    if UsersrulesruleMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesruleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

