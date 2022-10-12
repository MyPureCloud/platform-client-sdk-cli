package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentuserlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentuserlistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Learningassignmentuserlisting - List of users matching the learning module rule
type Learningassignmentuserlisting struct { 
    // Entities
    Entities []Learningassignmentuser `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total - The number of users matching search term
    Total int `json:"total"`


    // UnfilteredTotal - The total number of users
    UnfilteredTotal int `json:"unfilteredTotal"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentuserlisting) String() string {
     o.Entities = []Learningassignmentuser{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentuserlisting) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentuserlisting

    if LearningassignmentuserlistingMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentuserlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Learningassignmentuser `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        UnfilteredTotal int `json:"unfilteredTotal"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Learningassignmentuser{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

