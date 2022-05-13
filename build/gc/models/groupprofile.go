package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GroupprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GroupprofileDud struct { 
    Id string `json:"id"`


    


    


    DateModified time.Time `json:"dateModified"`


    


    SelfUri string `json:"selfUri"`

}

// Groupprofile
type Groupprofile struct { 
    


    // Name
    Name string `json:"name"`


    // State - The state of the user resource
    State string `json:"state"`


    


    // Version - The version of the group resource
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Groupprofile) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Groupprofile) MarshalJSON() ([]byte, error) {
    type Alias Groupprofile

    if GroupprofileMarshalled {
        return []byte("{}"), nil
    }
    GroupprofileMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        State string `json:"state"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

