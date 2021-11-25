package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsertimeoffrequestreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsertimeoffrequestreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Usertimeoffrequestreference
type Usertimeoffrequestreference struct { 
    // Id - The id of the time off request
    Id string `json:"id"`


    // User - The ID of the user to whom the time off request applies
    User Userreference `json:"user"`


    

}

// String returns a JSON representation of the model
func (o *Usertimeoffrequestreference) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usertimeoffrequestreference) MarshalJSON() ([]byte, error) {
    type Alias Usertimeoffrequestreference

    if UsertimeoffrequestreferenceMarshalled {
        return []byte("{}"), nil
    }
    UsertimeoffrequestreferenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        User Userreference `json:"user"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

