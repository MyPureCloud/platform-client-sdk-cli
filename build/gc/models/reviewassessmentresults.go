package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReviewassessmentresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReviewassessmentresultsDud struct { 
    


    

}

// Reviewassessmentresults
type Reviewassessmentresults struct { 
    // ByAssignees - If true, learning assignment results can be seen in detail by assignees
    ByAssignees bool `json:"byAssignees"`


    // ByViewers - If true, learning assignment results can be seen in detail by people who are eligible to view
    ByViewers bool `json:"byViewers"`

}

// String returns a JSON representation of the model
func (o *Reviewassessmentresults) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reviewassessmentresults) MarshalJSON() ([]byte, error) {
    type Alias Reviewassessmentresults

    if ReviewassessmentresultsMarshalled {
        return []byte("{}"), nil
    }
    ReviewassessmentresultsMarshalled = true

    return json.Marshal(&struct {
        
        ByAssignees bool `json:"byAssignees"`
        
        ByViewers bool `json:"byViewers"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

