package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamfullreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamfullreferenceDud struct { 
    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    MemberCount int `json:"memberCount"`


    SelfUri string `json:"selfUri"`

}

// Teamfullreference
type Teamfullreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name - The team name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - Team information.
    Description string `json:"description"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Teamfullreference) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamfullreference) MarshalJSON() ([]byte, error) {
    type Alias Teamfullreference

    if TeamfullreferenceMarshalled {
        return []byte("{}"), nil
    }
    TeamfullreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

