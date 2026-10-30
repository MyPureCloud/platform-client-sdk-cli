package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Orgv3sourcestagsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Orgv3sourcestagsresponseDud struct { 
    


    


    


    


    

}

// Orgv3sourcestagsresponse
type Orgv3sourcestagsresponse struct { 
    // Entities
    Entities []V3sourcetagresponse `json:"entities"`


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
func (o *Orgv3sourcestagsresponse) String() string {
     o.Entities = []V3sourcetagresponse{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Orgv3sourcestagsresponse) MarshalJSON() ([]byte, error) {
    type Alias Orgv3sourcestagsresponse

    if Orgv3sourcestagsresponseMarshalled {
        return []byte("{}"), nil
    }
    Orgv3sourcestagsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []V3sourcetagresponse `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []V3sourcetagresponse{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

