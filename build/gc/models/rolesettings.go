package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RolesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RolesettingsDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Rolesettings
type Rolesettings struct { 
    


    // BackfillEnabled - Boolean showing if organization is opted in or not to role backfills
    BackfillEnabled bool `json:"backfillEnabled"`


    // AuthorizationGrantDivisionAware - Boolean enabling enforcement of division aware for authorization grant add and delete
    AuthorizationGrantDivisionAware bool `json:"authorizationGrantDivisionAware"`


    

}

// String returns a JSON representation of the model
func (o *Rolesettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Rolesettings) MarshalJSON() ([]byte, error) {
    type Alias Rolesettings

    if RolesettingsMarshalled {
        return []byte("{}"), nil
    }
    RolesettingsMarshalled = true

    return json.Marshal(&struct {
        
        BackfillEnabled bool `json:"backfillEnabled"`
        
        AuthorizationGrantDivisionAware bool `json:"authorizationGrantDivisionAware"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

