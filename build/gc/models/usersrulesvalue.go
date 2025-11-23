package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesvalueDud struct { 
    


    

}

// Usersrulesvalue
type Usersrulesvalue struct { 
    // ContextId - The contextId for this group
    ContextId string `json:"contextId"`


    // Ids - The ids to select for this group
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Usersrulesvalue) String() string {
    
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesvalue) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesvalue

    if UsersrulesvalueMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesvalueMarshalled = true

    return json.Marshal(&struct {
        
        ContextId string `json:"contextId"`
        
        Ids []string `json:"ids"`
        *Alias
    }{

        


        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

