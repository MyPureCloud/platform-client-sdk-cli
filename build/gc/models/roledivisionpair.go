package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoledivisionpairMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoledivisionpairDud struct { 
    


    

}

// Roledivisionpair
type Roledivisionpair struct { 
    // RoleId - The ID of the role
    RoleId string `json:"roleId"`


    // DivisionId - The ID of the division
    DivisionId string `json:"divisionId"`

}

// String returns a JSON representation of the model
func (o *Roledivisionpair) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Roledivisionpair) MarshalJSON() ([]byte, error) {
    type Alias Roledivisionpair

    if RoledivisionpairMarshalled {
        return []byte("{}"), nil
    }
    RoledivisionpairMarshalled = true

    return json.Marshal(&struct { 
        RoleId string `json:"roleId"`
        
        DivisionId string `json:"divisionId"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

