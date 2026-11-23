package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueanomalyquerylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueanomalyquerylistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Queueanomalyquerylisting
type Queueanomalyquerylisting struct { 
    // Entities
    Entities []Queueanomalyresponse `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // Interval - Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // TimeZone - The time zone identifier for the organization, in IANA format
    TimeZone string `json:"timeZone"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Queueanomalyquerylisting) String() string {
     o.Entities = []Queueanomalyresponse{{}} 
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueanomalyquerylisting) MarshalJSON() ([]byte, error) {
    type Alias Queueanomalyquerylisting

    if QueueanomalyquerylistingMarshalled {
        return []byte("{}"), nil
    }
    QueueanomalyquerylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Queueanomalyresponse `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        Interval string `json:"interval"`
        
        TimeZone string `json:"timeZone"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Queueanomalyresponse{{}},
        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

