package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DisallowedentitylearningassignmentreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DisallowedentitylearningassignmentreferenceDud struct { 
    


    

}

// Disallowedentitylearningassignmentreference
type Disallowedentitylearningassignmentreference struct { 
    // ErrorCode - The error code associated with this disallowed entity
    ErrorCode string `json:"errorCode"`


    // Entity - The entity that was disallowed
    Entity Learningassignmentreference `json:"entity"`

}

// String returns a JSON representation of the model
func (o *Disallowedentitylearningassignmentreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Disallowedentitylearningassignmentreference) MarshalJSON() ([]byte, error) {
    type Alias Disallowedentitylearningassignmentreference

    if DisallowedentitylearningassignmentreferenceMarshalled {
        return []byte("{}"), nil
    }
    DisallowedentitylearningassignmentreferenceMarshalled = true

    return json.Marshal(&struct {
        
        ErrorCode string `json:"errorCode"`
        
        Entity Learningassignmentreference `json:"entity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

