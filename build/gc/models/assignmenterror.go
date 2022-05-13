package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignmenterrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignmenterrorDud struct { 
    


    

}

// Assignmenterror
type Assignmenterror struct { 
    // User - A user that is failed to be removed from the performance profile
    User Userreference `json:"user"`


    // Message - Error message from membership assignment
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Assignmenterror) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignmenterror) MarshalJSON() ([]byte, error) {
    type Alias Assignmenterror

    if AssignmenterrorMarshalled {
        return []byte("{}"), nil
    }
    AssignmenterrorMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

