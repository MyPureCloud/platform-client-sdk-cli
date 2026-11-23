package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalydetectionqueueanalysisMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalydetectionqueueanalysisDud struct { 
    


    


    


    


    

}

// Anomalydetectionqueueanalysis
type Anomalydetectionqueueanalysis struct { 
    // Metric - The name of the metric for which analysis was performed
    Metric string `json:"metric"`


    // MediaType - The media type that was analyzed
    MediaType string `json:"mediaType"`


    // DateLastUpdated - The timestamp when the model was last updated, in ISO-8601 format
    DateLastUpdated time.Time `json:"dateLastUpdated"`


    // AnalysisCode - When specified, indicates issues with the anomaly detection analysis
    AnalysisCode string `json:"analysisCode"`


    // TimeIntervalData - The time interval data with forecasted bounds. Not populated when analysisCode indicates an issue with the analysis
    TimeIntervalData []Anomalydetectiontimeintervaldata `json:"timeIntervalData"`

}

// String returns a JSON representation of the model
func (o *Anomalydetectionqueueanalysis) String() string {
    
    
    
    
     o.TimeIntervalData = []Anomalydetectiontimeintervaldata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalydetectionqueueanalysis) MarshalJSON() ([]byte, error) {
    type Alias Anomalydetectionqueueanalysis

    if AnomalydetectionqueueanalysisMarshalled {
        return []byte("{}"), nil
    }
    AnomalydetectionqueueanalysisMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        MediaType string `json:"mediaType"`
        
        DateLastUpdated time.Time `json:"dateLastUpdated"`
        
        AnalysisCode string `json:"analysisCode"`
        
        TimeIntervalData []Anomalydetectiontimeintervaldata `json:"timeIntervalData"`
        *Alias
    }{

        


        


        


        


        
        TimeIntervalData: []Anomalydetectiontimeintervaldata{{}},
        

        Alias: (*Alias)(u),
    })
}

