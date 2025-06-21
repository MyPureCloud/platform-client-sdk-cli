package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulemetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulemetadataDud struct { 
    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Buschedulemetadata
type Buschedulemetadata struct { 
    // Id - The ID of the schedule
    Id string `json:"id"`


    // WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // WeekCount - The number of weeks spanned by this schedule
    WeekCount int `json:"weekCount"`


    // Description - The description of this schedule
    Description string `json:"description"`


    // Published - Whether this schedule is published
    Published bool `json:"published"`


    // ShortTermForecast - The forecast used for this schedule, if applicable
    ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`


    // GenerationResults - Generation result summary for this schedule, if applicable
    GenerationResults Schedulegenerationresultsummary `json:"generationResults"`


    // Metadata - Version metadata for this schedule
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // ManagementUnits - High level per-management unit schedule metadata
    ManagementUnits []Bumanagementunitschedulesummary `json:"managementUnits"`


    

}

// String returns a JSON representation of the model
func (o *Buschedulemetadata) String() string {
    
    
    
    
    
    
    
    
     o.ManagementUnits = []Bumanagementunitschedulesummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulemetadata) MarshalJSON() ([]byte, error) {
    type Alias Buschedulemetadata

    if BuschedulemetadataMarshalled {
        return []byte("{}"), nil
    }
    BuschedulemetadataMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        WeekDate time.Time `json:"weekDate"`
        
        WeekCount int `json:"weekCount"`
        
        Description string `json:"description"`
        
        Published bool `json:"published"`
        
        ShortTermForecast Bushorttermforecastreference `json:"shortTermForecast"`
        
        GenerationResults Schedulegenerationresultsummary `json:"generationResults"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        ManagementUnits []Bumanagementunitschedulesummary `json:"managementUnits"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        ManagementUnits: []Bumanagementunitschedulesummary{{}},
        


        

        Alias: (*Alias)(u),
    })
}

