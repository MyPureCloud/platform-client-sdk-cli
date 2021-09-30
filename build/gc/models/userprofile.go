package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserprofileDud struct { 
    Id string `json:"id"`


    


    


    DateModified time.Time `json:"dateModified"`


    


    Expands Userexpands `json:"expands"`


    SelfUri string `json:"selfUri"`

}

// Userprofile
type Userprofile struct { 
    


    // Name
    Name string `json:"name"`


    // State - The state of the user resource
    State string `json:"state"`


    


    // Version - The version of the group resource
    Version int `json:"version"`


    


    

}

// String returns a JSON representation of the model
func (o *Userprofile) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userprofile) MarshalJSON() ([]byte, error) {
    type Alias Userprofile

    if UserprofileMarshalled {
        return []byte("{}"), nil
    }
    UserprofileMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        State string `json:"state"`
        
        
        
        Version int `json:"version"`
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

