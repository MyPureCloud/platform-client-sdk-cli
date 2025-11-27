package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReprocessjobentitylistingresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReprocessjobentitylistingresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Reprocessjobentitylistingresponse
type Reprocessjobentitylistingresponse struct { 
    // Entities
    Entities []Reprocessjobresponse `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // QueueTotal - The total number of queued jobs.
    QueueTotal int `json:"queueTotal"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Reprocessjobentitylistingresponse) String() string {
     o.Entities = []Reprocessjobresponse{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reprocessjobentitylistingresponse) MarshalJSON() ([]byte, error) {
    type Alias Reprocessjobentitylistingresponse

    if ReprocessjobentitylistingresponseMarshalled {
        return []byte("{}"), nil
    }
    ReprocessjobentitylistingresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Reprocessjobresponse `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        QueueTotal int `json:"queueTotal"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Reprocessjobresponse{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

