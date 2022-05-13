package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoledivisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoledivisionDud struct { 
    


    

}

// Roledivision
type Roledivision struct { 
    // RoleId - Role to be associated with the given division which forms a grant
    RoleId string `json:"roleId"`


    // DivisionId - Division associated with the given role which forms a grant
    DivisionId string `json:"divisionId"`

}

// String returns a JSON representation of the model
func (o *Roledivision) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Roledivision) MarshalJSON() ([]byte, error) {
    type Alias Roledivision

    if RoledivisionMarshalled {
        return []byte("{}"), nil
    }
    RoledivisionMarshalled = true

    return json.Marshal(&struct {
        
        RoleId string `json:"roleId"`
        
        DivisionId string `json:"divisionId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

