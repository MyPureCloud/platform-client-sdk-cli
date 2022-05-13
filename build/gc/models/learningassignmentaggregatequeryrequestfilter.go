package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentaggregatequeryrequestfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentaggregatequeryrequestfilterDud struct { 
    


    

}

// Learningassignmentaggregatequeryrequestfilter
type Learningassignmentaggregatequeryrequestfilter struct { 
    // VarType - The logic used to combine the clauses
    VarType string `json:"type"`


    // Clauses - The list of clauses used to filter the data. Note that clauses must filter by attendeeId and a maximum of 100 user IDs are allowed
    Clauses []Learningassignmentaggregatequeryrequestclause `json:"clauses"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestfilter) String() string {
    
     o.Clauses = []Learningassignmentaggregatequeryrequestclause{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentaggregatequeryrequestfilter) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentaggregatequeryrequestfilter

    if LearningassignmentaggregatequeryrequestfilterMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentaggregatequeryrequestfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Learningassignmentaggregatequeryrequestclause `json:"clauses"`
        *Alias
    }{

        


        
        Clauses: []Learningassignmentaggregatequeryrequestclause{{}},
        

        Alias: (*Alias)(u),
    })
}

