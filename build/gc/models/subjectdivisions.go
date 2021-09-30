package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SubjectdivisionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SubjectdivisionsDud struct { 
    


    

}

// Subjectdivisions
type Subjectdivisions struct { 
    // SubjectIds - A collection of subject IDs to associate with the given divisions
    SubjectIds []string `json:"subjectIds"`


    // DivisionIds - A collection of division IDs to associate with the given subjects
    DivisionIds []string `json:"divisionIds"`

}

// String returns a JSON representation of the model
func (o *Subjectdivisions) String() string {
    
    
     o.SubjectIds = []string{""} 
    
    
    
     o.DivisionIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Subjectdivisions) MarshalJSON() ([]byte, error) {
    type Alias Subjectdivisions

    if SubjectdivisionsMarshalled {
        return []byte("{}"), nil
    }
    SubjectdivisionsMarshalled = true

    return json.Marshal(&struct { 
        SubjectIds []string `json:"subjectIds"`
        
        DivisionIds []string `json:"divisionIds"`
        
        *Alias
    }{
        

        
        SubjectIds: []string{""},
        

        

        
        DivisionIds: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

