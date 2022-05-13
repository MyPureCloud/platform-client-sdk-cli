package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmoduleruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmoduleruleDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Learningmodulerule
type Learningmodulerule struct { 
    


    // IsActive - If true, rule is active
    IsActive bool `json:"isActive"`


    // Parts - The parts of a learning module rule
    Parts []Learningmoduleruleparts `json:"parts"`


    

}

// String returns a JSON representation of the model
func (o *Learningmodulerule) String() string {
    
     o.Parts = []Learningmoduleruleparts{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulerule) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulerule

    if LearningmoduleruleMarshalled {
        return []byte("{}"), nil
    }
    LearningmoduleruleMarshalled = true

    return json.Marshal(&struct {
        
        IsActive bool `json:"isActive"`
        
        Parts []Learningmoduleruleparts `json:"parts"`
        *Alias
    }{

        


        


        
        Parts: []Learningmoduleruleparts{{}},
        


        

        Alias: (*Alias)(u),
    })
}

