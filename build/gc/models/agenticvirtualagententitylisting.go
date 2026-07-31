package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagententitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagententitylistingDud struct { 
    


    


    


    


    

}

// Agenticvirtualagententitylisting
type Agenticvirtualagententitylisting struct { 
    // Entities
    Entities []Agenticvirtualagent `json:"entities"`


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
func (o *Agenticvirtualagententitylisting) String() string {
     o.Entities = []Agenticvirtualagent{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagententitylisting) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagententitylisting

    if AgenticvirtualagententitylistingMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagententitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Agenticvirtualagent `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Agenticvirtualagent{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

