package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulelistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulelistDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Learningmodulelist - Learning module list
type Learningmodulelist struct { 
    // Entities
    Entities []Learningmodule `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // TotalLegacyRules - The total number of unmigrated rules
    TotalLegacyRules int `json:"totalLegacyRules"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Learningmodulelist) String() string {
     o.Entities = []Learningmodule{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulelist) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulelist

    if LearningmodulelistMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulelistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Learningmodule `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        TotalLegacyRules int `json:"totalLegacyRules"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Learningmodule{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

