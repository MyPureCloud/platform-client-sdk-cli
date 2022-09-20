package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TestmatcheseventoperationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TestmatcheseventoperationDud struct { 
    


    


    


    

}

// Testmatcheseventoperation - Results from evaluating matching criteria against test input
type Testmatcheseventoperation struct { 
    // Name - The name of the processing step
    Name string `json:"name"`


    // Step - The number of the processing step
    Step int `json:"step"`


    // MatchedTriggers - Triggers that matched
    MatchedTriggers []Testmodetrigger `json:"matchedTriggers"`


    // UnmatchedTriggers - Triggers that did not match
    UnmatchedTriggers []Testmodetrigger `json:"unmatchedTriggers"`

}

// String returns a JSON representation of the model
func (o *Testmatcheseventoperation) String() string {
    
    
     o.MatchedTriggers = []Testmodetrigger{{}} 
     o.UnmatchedTriggers = []Testmodetrigger{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testmatcheseventoperation) MarshalJSON() ([]byte, error) {
    type Alias Testmatcheseventoperation

    if TestmatcheseventoperationMarshalled {
        return []byte("{}"), nil
    }
    TestmatcheseventoperationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Step int `json:"step"`
        
        MatchedTriggers []Testmodetrigger `json:"matchedTriggers"`
        
        UnmatchedTriggers []Testmodetrigger `json:"unmatchedTriggers"`
        *Alias
    }{

        


        


        
        MatchedTriggers: []Testmodetrigger{{}},
        


        
        UnmatchedTriggers: []Testmodetrigger{{}},
        

        Alias: (*Alias)(u),
    })
}

