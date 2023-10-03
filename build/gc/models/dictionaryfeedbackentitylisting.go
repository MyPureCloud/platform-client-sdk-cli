package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DictionaryfeedbackentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DictionaryfeedbackentitylistingDud struct { 
    


    


    


    


    


    

}

// Dictionaryfeedbackentitylisting
type Dictionaryfeedbackentitylisting struct { 
    // Entities
    Entities []Listeddictionaryfeedback `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageCount
    PageCount int `json:"pageCount"`


    // Total
    Total int `json:"total"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`

}

// String returns a JSON representation of the model
func (o *Dictionaryfeedbackentitylisting) String() string {
     o.Entities = []Listeddictionaryfeedback{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dictionaryfeedbackentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Dictionaryfeedbackentitylisting

    if DictionaryfeedbackentitylistingMarshalled {
        return []byte("{}"), nil
    }
    DictionaryfeedbackentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Listeddictionaryfeedback `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageCount int `json:"pageCount"`
        
        Total int `json:"total"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        *Alias
    }{

        
        Entities: []Listeddictionaryfeedback{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

