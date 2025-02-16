package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SubjectMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SubjectDud struct { 
    


    

}

// Subject
type Subject struct { 
    // VarType
    VarType string `json:"type"`


    // Id
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Subject) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Subject) MarshalJSON() ([]byte, error) {
    type Alias Subject

    if SubjectMarshalled {
        return []byte("{}"), nil
    }
    SubjectMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

