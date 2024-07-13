package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersearchrulepartMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersearchrulepartDud struct { 
    


    


    

}

// Usersearchrulepart
type Usersearchrulepart struct { 
    // Operation - The operation for this part
    Operation string `json:"operation"`


    // Selector - The type of item being selected by this part
    Selector string `json:"selector"`


    // Ids - The ids for the selector; the results of these are ORed together
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Usersearchrulepart) String() string {
    
    
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersearchrulepart) MarshalJSON() ([]byte, error) {
    type Alias Usersearchrulepart

    if UsersearchrulepartMarshalled {
        return []byte("{}"), nil
    }
    UsersearchrulepartMarshalled = true

    return json.Marshal(&struct {
        
        Operation string `json:"operation"`
        
        Selector string `json:"selector"`
        
        Ids []string `json:"ids"`
        *Alias
    }{

        


        


        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

