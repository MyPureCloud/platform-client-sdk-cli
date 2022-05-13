package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SubjectdivisiongrantsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SubjectdivisiongrantsDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Subjectdivisiongrants
type Subjectdivisiongrants struct { 
    


    // Name
    Name string `json:"name"`


    // Divisions
    Divisions []Division `json:"divisions"`


    // VarType
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Subjectdivisiongrants) String() string {
    
     o.Divisions = []Division{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Subjectdivisiongrants) MarshalJSON() ([]byte, error) {
    type Alias Subjectdivisiongrants

    if SubjectdivisiongrantsMarshalled {
        return []byte("{}"), nil
    }
    SubjectdivisiongrantsMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Divisions []Division `json:"divisions"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        
        Divisions: []Division{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

