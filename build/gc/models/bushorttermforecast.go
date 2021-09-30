package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BushorttermforecastMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BushorttermforecastDud struct { 
    Id string `json:"id"`


    


    


    


    


    Legacy bool `json:"legacy"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Bushorttermforecast
type Bushorttermforecast struct { 
    


    // WeekDate - The start week date of this forecast in yyyy-MM-dd.  Must fall on the start day of week for the associated business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // WeekCount - The number of weeks this forecast covers
    WeekCount int `json:"weekCount"`


    // CreationMethod - The method by which this forecast was created
    CreationMethod string `json:"creationMethod"`


    // Description - The description of this forecast
    Description string `json:"description"`


    


    // Metadata - Metadata for this forecast
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // CanUseForScheduling - Whether this forecast can be used for scheduling
    CanUseForScheduling bool `json:"canUseForScheduling"`


    // ReferenceStartDate - The reference start date for interval-based data for this forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReferenceStartDate time.Time `json:"referenceStartDate"`


    // SourceDays - The source day pointers for this forecast
    SourceDays []Forecastsourcedaypointer `json:"sourceDays"`


    // Modifications - Any manual modifications applied to this forecast
    Modifications []Buforecastmodification `json:"modifications"`


    // GenerationResults - Generation result metadata
    GenerationResults Buforecastgenerationresult `json:"generationResults"`


    // TimeZone - The time zone for this forecast
    TimeZone string `json:"timeZone"`


    // PlanningGroupsVersion - The version of the planning groups that was used for this forecast
    PlanningGroupsVersion int `json:"planningGroupsVersion"`


    // PlanningGroups - A snapshot of the planning groups used for this forecast as of the version number indicated
    PlanningGroups Forecastplanninggroupsresponse `json:"planningGroups"`


    

}

// String returns a JSON representation of the model
func (o *Bushorttermforecast) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.SourceDays = []Forecastsourcedaypointer{{}} 
    
    
    
     o.Modifications = []Buforecastmodification{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bushorttermforecast) MarshalJSON() ([]byte, error) {
    type Alias Bushorttermforecast

    if BushorttermforecastMarshalled {
        return []byte("{}"), nil
    }
    BushorttermforecastMarshalled = true

    return json.Marshal(&struct { 
        
        
        WeekDate time.Time `json:"weekDate"`
        
        WeekCount int `json:"weekCount"`
        
        CreationMethod string `json:"creationMethod"`
        
        Description string `json:"description"`
        
        
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        CanUseForScheduling bool `json:"canUseForScheduling"`
        
        ReferenceStartDate time.Time `json:"referenceStartDate"`
        
        SourceDays []Forecastsourcedaypointer `json:"sourceDays"`
        
        Modifications []Buforecastmodification `json:"modifications"`
        
        GenerationResults Buforecastgenerationresult `json:"generationResults"`
        
        TimeZone string `json:"timeZone"`
        
        PlanningGroupsVersion int `json:"planningGroupsVersion"`
        
        PlanningGroups Forecastplanninggroupsresponse `json:"planningGroups"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        SourceDays: []Forecastsourcedaypointer{{}},
        

        

        
        Modifications: []Buforecastmodification{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

