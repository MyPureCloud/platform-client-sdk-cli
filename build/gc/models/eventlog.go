package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventlogMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventlogDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Eventlog
type Eventlog struct { 
    


    // Name
    Name string `json:"name"`


    // ErrorEntity
    ErrorEntity Domainentityref `json:"errorEntity"`


    // RelatedEntity
    RelatedEntity Domainentityref `json:"relatedEntity"`


    // Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    // Level
    Level string `json:"level"`


    // Category
    Category string `json:"category"`


    // CorrelationId
    CorrelationId string `json:"correlationId"`


    // EventMessage
    EventMessage Eventmessage `json:"eventMessage"`


    

}

// String returns a JSON representation of the model
func (o *Eventlog) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventlog) MarshalJSON() ([]byte, error) {
    type Alias Eventlog

    if EventlogMarshalled {
        return []byte("{}"), nil
    }
    EventlogMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ErrorEntity Domainentityref `json:"errorEntity"`
        
        RelatedEntity Domainentityref `json:"relatedEntity"`
        
        Timestamp time.Time `json:"timestamp"`
        
        Level string `json:"level"`
        
        Category string `json:"category"`
        
        CorrelationId string `json:"correlationId"`
        
        EventMessage Eventmessage `json:"eventMessage"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

