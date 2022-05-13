package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzdivisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzdivisionDud struct { 
    Id string `json:"id"`


    


    


    HomeDivision bool `json:"homeDivision"`


    ObjectCounts map[string]int `json:"objectCounts"`


    SelfUri string `json:"selfUri"`

}

// Authzdivision
type Authzdivision struct { 
    


    // Name
    Name string `json:"name"`


    // Description - A helpful description for the division.
    Description string `json:"description"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Authzdivision) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzdivision) MarshalJSON() ([]byte, error) {
    type Alias Authzdivision

    if AuthzdivisionMarshalled {
        return []byte("{}"), nil
    }
    AuthzdivisionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

