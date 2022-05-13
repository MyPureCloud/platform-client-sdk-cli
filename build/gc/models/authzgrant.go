package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzgrantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzgrantDud struct { 
    


    


    


    

}

// Authzgrant
type Authzgrant struct { 
    // SubjectId
    SubjectId string `json:"subjectId"`


    // Division
    Division Authzdivision `json:"division"`


    // Role
    Role Authzgrantrole `json:"role"`


    // GrantMadeAt - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    GrantMadeAt time.Time `json:"grantMadeAt"`

}

// String returns a JSON representation of the model
func (o *Authzgrant) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzgrant) MarshalJSON() ([]byte, error) {
    type Alias Authzgrant

    if AuthzgrantMarshalled {
        return []byte("{}"), nil
    }
    AuthzgrantMarshalled = true

    return json.Marshal(&struct {
        
        SubjectId string `json:"subjectId"`
        
        Division Authzdivision `json:"division"`
        
        Role Authzgrantrole `json:"role"`
        
        GrantMadeAt time.Time `json:"grantMadeAt"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

