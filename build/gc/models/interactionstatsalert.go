package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InteractionstatsalertMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InteractionstatsalertDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    Dimension string `json:"dimension"`


    DimensionValue string `json:"dimensionValue"`


    Metric string `json:"metric"`


    MediaType string `json:"mediaType"`


    NumericRange string `json:"numericRange"`


    Statistic string `json:"statistic"`


    Value float64 `json:"value"`


    RuleId string `json:"ruleId"`


    


    StartDate time.Time `json:"startDate"`


    EndDate time.Time `json:"endDate"`


    NotificationUsers []User `json:"notificationUsers"`


    AlertTypes []string `json:"alertTypes"`


    


    SelfUri string `json:"selfUri"`

}

// Interactionstatsalert
type Interactionstatsalert struct { 
    


    


    


    


    


    


    


    


    


    


    // Unread - Indicates if the alert has been read.
    Unread bool `json:"unread"`


    


    


    


    


    // RuleUri
    RuleUri string `json:"ruleUri"`


    

}

// String returns a JSON representation of the model
func (o *Interactionstatsalert) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Interactionstatsalert) MarshalJSON() ([]byte, error) {
    type Alias Interactionstatsalert

    if InteractionstatsalertMarshalled {
        return []byte("{}"), nil
    }
    InteractionstatsalertMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        Unread bool `json:"unread"`
        
        
        
        
        
        
        
        
        
        RuleUri string `json:"ruleUri"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

