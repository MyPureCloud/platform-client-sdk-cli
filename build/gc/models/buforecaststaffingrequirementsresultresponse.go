package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuforecaststaffingrequirementsresultresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuforecaststaffingrequirementsresultresponseDud struct { 
    


    


    


    


    


    


    

}

// Buforecaststaffingrequirementsresultresponse
type Buforecaststaffingrequirementsresultresponse struct { 
    // BusinessUnitId - The ID of the business unit to which the forecast staffing requirements belongs
    BusinessUnitId string `json:"businessUnitId"`


    // Forecast - The forecast reference
    Forecast Bushorttermforecastreference `json:"forecast"`


    // ReferenceStartDate - The reference start date for interval-based data for this forecast as an ISO-8601 string
    ReferenceStartDate time.Time `json:"referenceStartDate"`


    // WeekCount - The number of weeks in this forecast
    WeekCount int `json:"weekCount"`


    // IntervalLengthMinutes - The period/interval in minutes for which to aggregate the data
    IntervalLengthMinutes int `json:"intervalLengthMinutes"`


    // State - The state of the staffing requirements generation
    State string `json:"state"`


    // Results - The forecast staffing requirement results, Will be populated when state == 'Complete'
    Results []Buforecaststaffingrequirementsresult `json:"results"`

}

// String returns a JSON representation of the model
func (o *Buforecaststaffingrequirementsresultresponse) String() string {
    
    
    
    
    
    
     o.Results = []Buforecaststaffingrequirementsresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buforecaststaffingrequirementsresultresponse) MarshalJSON() ([]byte, error) {
    type Alias Buforecaststaffingrequirementsresultresponse

    if BuforecaststaffingrequirementsresultresponseMarshalled {
        return []byte("{}"), nil
    }
    BuforecaststaffingrequirementsresultresponseMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnitId string `json:"businessUnitId"`
        
        Forecast Bushorttermforecastreference `json:"forecast"`
        
        ReferenceStartDate time.Time `json:"referenceStartDate"`
        
        WeekCount int `json:"weekCount"`
        
        IntervalLengthMinutes int `json:"intervalLengthMinutes"`
        
        State string `json:"state"`
        
        Results []Buforecaststaffingrequirementsresult `json:"results"`
        *Alias
    }{

        


        


        


        


        


        


        
        Results: []Buforecaststaffingrequirementsresult{{}},
        

        Alias: (*Alias)(u),
    })
}

