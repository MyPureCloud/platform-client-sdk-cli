package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DisallowedentitylearningassignmentitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DisallowedentitylearningassignmentitemDud struct { 
    


    

}

// Disallowedentitylearningassignmentitem
type Disallowedentitylearningassignmentitem struct { 
    // ErrorCode - The error code associated with this disallowed entity
    ErrorCode string `json:"errorCode"`


    // Entity - The entity that was disallowed
    Entity Learningassignmentitem `json:"entity"`

}

// String returns a JSON representation of the model
func (o *Disallowedentitylearningassignmentitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Disallowedentitylearningassignmentitem) MarshalJSON() ([]byte, error) {
    type Alias Disallowedentitylearningassignmentitem

    if DisallowedentitylearningassignmentitemMarshalled {
        return []byte("{}"), nil
    }
    DisallowedentitylearningassignmentitemMarshalled = true

    return json.Marshal(&struct {
        
        ErrorCode string `json:"errorCode"`
        
        Entity Learningassignmentitem `json:"entity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

