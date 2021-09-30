package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekschedulelistitemresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekschedulelistitemresponseDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`


    


    


    


    


    


    

}

// Weekschedulelistitemresponse
type Weekschedulelistitemresponse struct { 
    


    


    // WeekDate - First day of this week schedule in yyyy-MM-dd format
    WeekDate string `json:"weekDate"`


    // Description - Description of the week schedule
    Description string `json:"description"`


    // Published - Whether the week schedule is published
    Published bool `json:"published"`


    // GenerationResults - Summary of the results from the schedule run
    GenerationResults Weekschedulegenerationresult `json:"generationResults"`


    // ShortTermForecast - Short term forecast associated with this schedule
    ShortTermForecast Shorttermforecastreference `json:"shortTermForecast"`


    // Metadata - Version metadata for this work plan
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Weekschedulelistitemresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekschedulelistitemresponse) MarshalJSON() ([]byte, error) {
    type Alias Weekschedulelistitemresponse

    if WeekschedulelistitemresponseMarshalled {
        return []byte("{}"), nil
    }
    WeekschedulelistitemresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        WeekDate string `json:"weekDate"`
        
        Description string `json:"description"`
        
        Published bool `json:"published"`
        
        GenerationResults Weekschedulegenerationresult `json:"generationResults"`
        
        ShortTermForecast Shorttermforecastreference `json:"shortTermForecast"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

