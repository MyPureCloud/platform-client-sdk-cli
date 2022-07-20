package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalmetricdataunprocesseditemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalmetricdataunprocesseditemDud struct { 
    


    


    


    


    


    


    


    

}

// Externalmetricdataunprocesseditem
type Externalmetricdataunprocesseditem struct { 
    // UserId - The user ID. Must provide either userId or userEmail, but not both.
    UserId string `json:"userId"`


    // UserEmail - The user main email used in user's GenesysCloud account. Must provide either userId or userEmail, but not both.
    UserEmail string `json:"userEmail"`


    // MetricId - The ID of the external metric definition
    MetricId string `json:"metricId"`


    // DateOccurred - The date of the metric data. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateOccurred time.Time `json:"dateOccurred"`


    // Value - The value of the metric data. When value is null, the metric data will be deleted.
    Value float64 `json:"value"`


    // Count - The number of data points. The default value is 1.
    Count int `json:"count"`


    // Message - The error message
    Message string `json:"message"`


    // Code - The error code
    Code string `json:"code"`

}

// String returns a JSON representation of the model
func (o *Externalmetricdataunprocesseditem) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalmetricdataunprocesseditem) MarshalJSON() ([]byte, error) {
    type Alias Externalmetricdataunprocesseditem

    if ExternalmetricdataunprocesseditemMarshalled {
        return []byte("{}"), nil
    }
    ExternalmetricdataunprocesseditemMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        UserEmail string `json:"userEmail"`
        
        MetricId string `json:"metricId"`
        
        DateOccurred time.Time `json:"dateOccurred"`
        
        Value float64 `json:"value"`
        
        Count int `json:"count"`
        
        Message string `json:"message"`
        
        Code string `json:"code"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

