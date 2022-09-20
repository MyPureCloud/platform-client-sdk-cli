package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentreassignMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentreassignDud struct { 
    


    

}

// Learningassignmentreassign
type Learningassignmentreassign struct { 
    // RecommendedCompletionDate - The recommended completion date of assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RecommendedCompletionDate time.Time `json:"recommendedCompletionDate"`


    // LengthInMinutes - The length in minutes of assignment
    LengthInMinutes int `json:"lengthInMinutes"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentreassign) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentreassign) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentreassign

    if LearningassignmentreassignMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentreassignMarshalled = true

    return json.Marshal(&struct {
        
        RecommendedCompletionDate time.Time `json:"recommendedCompletionDate"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

