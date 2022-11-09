package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalmetricdataitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalmetricdataitemDud struct { 
    


    


    


    


    


    


    

}

// Externalmetricdataitem
type Externalmetricdataitem struct { 
    // UserId - The user ID. Must provide either userId or userEmail, but not both.
    UserId string `json:"userId"`


    // UserEmail - The user main email used in user's GenesysCloud account. Must provide either userId or userEmail, but not both.
    UserEmail string `json:"userEmail"`


    // MetricId - The ID of the external metric definition
    MetricId string `json:"metricId"`


    // DateOccurred - The date of the metric data. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateOccurred time.Time `json:"dateOccurred"`


    // Value - The value of the metric data. When value is null, the metric data will be deleted.
    Value float32 `json:"value"`


    // Count - The number of data points. The default value is 0 when type is Cumulative and the metric data already exists, otherwise 1. When total count reaches 0, the metric data will be deleted.
    Count int `json:"count"`


    // VarType - The type of the metric data. The default value is Total.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Externalmetricdataitem) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalmetricdataitem) MarshalJSON() ([]byte, error) {
    type Alias Externalmetricdataitem

    if ExternalmetricdataitemMarshalled {
        return []byte("{}"), nil
    }
    ExternalmetricdataitemMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        UserEmail string `json:"userEmail"`
        
        MetricId string `json:"metricId"`
        
        DateOccurred time.Time `json:"dateOccurred"`
        
        Value float32 `json:"value"`
        
        Count int `json:"count"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

