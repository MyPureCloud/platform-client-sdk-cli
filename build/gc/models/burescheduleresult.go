package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BurescheduleresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BurescheduleresultDud struct { 
    


    


    


    


    

}

// Burescheduleresult
type Burescheduleresult struct { 
    // GenerationResults - The generation results.  Note the result will always be delivered via the downloadUrl; however the schema is included for documentation
    GenerationResults Schedulegenerationresult `json:"generationResults"`


    // GenerationResultsDownloadUrl - The download URL from which to fetch the generation results for the rescheduling run
    GenerationResultsDownloadUrl string `json:"generationResultsDownloadUrl"`


    // HeadcountForecast - The headcount forecast.  Note the result will always be delivered via the downloadUrl; however the schema is included for documentation
    HeadcountForecast Buheadcountforecast `json:"headcountForecast"`


    // HeadcountForecastDownloadUrl - The download URL from which to fetch the headcount forecast for the rescheduling run
    HeadcountForecastDownloadUrl string `json:"headcountForecastDownloadUrl"`


    // AgentSchedules - List of download links for agent schedules produced by the rescheduling run
    AgentSchedules []Burescheduleagentscheduleresult `json:"agentSchedules"`

}

// String returns a JSON representation of the model
func (o *Burescheduleresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AgentSchedules = []Burescheduleagentscheduleresult{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Burescheduleresult) MarshalJSON() ([]byte, error) {
    type Alias Burescheduleresult

    if BurescheduleresultMarshalled {
        return []byte("{}"), nil
    }
    BurescheduleresultMarshalled = true

    return json.Marshal(&struct { 
        GenerationResults Schedulegenerationresult `json:"generationResults"`
        
        GenerationResultsDownloadUrl string `json:"generationResultsDownloadUrl"`
        
        HeadcountForecast Buheadcountforecast `json:"headcountForecast"`
        
        HeadcountForecastDownloadUrl string `json:"headcountForecastDownloadUrl"`
        
        AgentSchedules []Burescheduleagentscheduleresult `json:"agentSchedules"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        AgentSchedules: []Burescheduleagentscheduleresult{{}},
        

        
        Alias: (*Alias)(u),
    })
}

