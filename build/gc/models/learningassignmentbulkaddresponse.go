package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentbulkaddresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentbulkaddresponseDud struct { 
    


    

}

// Learningassignmentbulkaddresponse
type Learningassignmentbulkaddresponse struct { 
    // Entities - The learning assignments that were assigned correctly
    Entities []Learningassignment `json:"entities"`


    // DisallowedEntities - The items that were not allowed to be assigned
    DisallowedEntities []Disallowedentitylearningassignmentitem `json:"disallowedEntities"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentbulkaddresponse) String() string {
     o.Entities = []Learningassignment{{}} 
     o.DisallowedEntities = []Disallowedentitylearningassignmentitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentbulkaddresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentbulkaddresponse

    if LearningassignmentbulkaddresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentbulkaddresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Learningassignment `json:"entities"`
        
        DisallowedEntities []Disallowedentitylearningassignmentitem `json:"disallowedEntities"`
        *Alias
    }{

        
        Entities: []Learningassignment{{}},
        


        
        DisallowedEntities: []Disallowedentitylearningassignmentitem{{}},
        

        Alias: (*Alias)(u),
    })
}

