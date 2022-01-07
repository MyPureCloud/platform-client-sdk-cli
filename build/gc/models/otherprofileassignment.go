package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OtherprofileassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OtherprofileassignmentDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Otherprofileassignment
type Otherprofileassignment struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // CurrentProfile - The current performance profile that this user belongs to
    CurrentProfile Domainentityref `json:"currentProfile"`


    

}

// String returns a JSON representation of the model
func (o *Otherprofileassignment) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Otherprofileassignment) MarshalJSON() ([]byte, error) {
    type Alias Otherprofileassignment

    if OtherprofileassignmentMarshalled {
        return []byte("{}"), nil
    }
    OtherprofileassignmentMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        CurrentProfile Domainentityref `json:"currentProfile"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

