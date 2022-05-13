package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentbulkremoveresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentbulkremoveresponseDud struct { 
    


    

}

// Learningassignmentbulkremoveresponse
type Learningassignmentbulkremoveresponse struct { 
    // Entities - The learning assignments that were removed successfully
    Entities []Learningassignmententity `json:"entities"`


    // DisallowedEntities - The learning assignments that were not removed due to missing permissions
    DisallowedEntities []Disallowedentitylearningassignmentreference `json:"disallowedEntities"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentbulkremoveresponse) String() string {
     o.Entities = []Learningassignmententity{{}} 
     o.DisallowedEntities = []Disallowedentitylearningassignmentreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentbulkremoveresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentbulkremoveresponse

    if LearningassignmentbulkremoveresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentbulkremoveresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Learningassignmententity `json:"entities"`
        
        DisallowedEntities []Disallowedentitylearningassignmentreference `json:"disallowedEntities"`
        *Alias
    }{

        
        Entities: []Learningassignmententity{{}},
        


        
        DisallowedEntities: []Disallowedentitylearningassignmentreference{{}},
        

        Alias: (*Alias)(u),
    })
}

