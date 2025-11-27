package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecognitionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecognitionsDud struct { 
    


    


    


    


    

}

// Recognitions
type Recognitions struct { 
    // Entities
    Entities []Recognition `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Recognitions) String() string {
     o.Entities = []Recognition{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recognitions) MarshalJSON() ([]byte, error) {
    type Alias Recognitions

    if RecognitionsMarshalled {
        return []byte("{}"), nil
    }
    RecognitionsMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Recognition `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Recognition{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

