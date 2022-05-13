package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InteractionstatsruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InteractionstatsruleDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    InAlarm bool `json:"inAlarm"`


    


    


    SelfUri string `json:"selfUri"`

}

// Interactionstatsrule
type Interactionstatsrule struct { 
    


    // Name - Name of the rule
    Name string `json:"name"`


    // Dimension - The dimension of concern.
    Dimension string `json:"dimension"`


    // DimensionValue - The value of the dimension.
    DimensionValue string `json:"dimensionValue"`


    // Metric - The metric to be assessed.
    Metric string `json:"metric"`


    // MediaType - The media type.
    MediaType string `json:"mediaType"`


    // NumericRange - The comparison descriptor used against the metric's value.
    NumericRange string `json:"numericRange"`


    // Statistic - The statistic of concern for the metric.
    Statistic string `json:"statistic"`


    // Value - The threshold value.
    Value float64 `json:"value"`


    // Enabled - Indicates if the rule is enabled.
    Enabled bool `json:"enabled"`


    


    // NotificationUsers - The ids of users who will be notified of alarm state change.
    NotificationUsers []User `json:"notificationUsers"`


    // AlertTypes - A collection of notification methods.
    AlertTypes []string `json:"alertTypes"`


    

}

// String returns a JSON representation of the model
func (o *Interactionstatsrule) String() string {
    
    
    
    
    
    
    
    
    
     o.NotificationUsers = []User{{}} 
     o.AlertTypes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Interactionstatsrule) MarshalJSON() ([]byte, error) {
    type Alias Interactionstatsrule

    if InteractionstatsruleMarshalled {
        return []byte("{}"), nil
    }
    InteractionstatsruleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Dimension string `json:"dimension"`
        
        DimensionValue string `json:"dimensionValue"`
        
        Metric string `json:"metric"`
        
        MediaType string `json:"mediaType"`
        
        NumericRange string `json:"numericRange"`
        
        Statistic string `json:"statistic"`
        
        Value float64 `json:"value"`
        
        Enabled bool `json:"enabled"`
        
        NotificationUsers []User `json:"notificationUsers"`
        
        AlertTypes []string `json:"alertTypes"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        NotificationUsers: []User{{}},
        


        
        AlertTypes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

