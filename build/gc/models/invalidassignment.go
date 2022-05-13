package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InvalidassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InvalidassignmentDud struct { 
    


    

}

// Invalidassignment
type Invalidassignment struct { 
    // User - Invalid user for validation
    User Userreference `json:"user"`


    // Message - The reason for the invalid input for validation
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Invalidassignment) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Invalidassignment) MarshalJSON() ([]byte, error) {
    type Alias Invalidassignment

    if InvalidassignmentMarshalled {
        return []byte("{}"), nil
    }
    InvalidassignmentMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

