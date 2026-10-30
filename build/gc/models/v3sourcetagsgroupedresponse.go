package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcetagsgroupedresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcetagsgroupedresponseDud struct { 
    


    


    


    


    

}

// V3sourcetagsgroupedresponse
type V3sourcetagsgroupedresponse struct { 
    // Entities
    Entities []V3sourcetagsgroupedentry `json:"entities"`


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
func (o *V3sourcetagsgroupedresponse) String() string {
     o.Entities = []V3sourcetagsgroupedentry{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcetagsgroupedresponse) MarshalJSON() ([]byte, error) {
    type Alias V3sourcetagsgroupedresponse

    if V3sourcetagsgroupedresponseMarshalled {
        return []byte("{}"), nil
    }
    V3sourcetagsgroupedresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []V3sourcetagsgroupedentry `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []V3sourcetagsgroupedentry{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

