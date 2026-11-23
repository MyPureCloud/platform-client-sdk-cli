package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueanomalyexportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueanomalyexportDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Queueanomalyexport
type Queueanomalyexport struct { 
    // Id - The ID of the queue anomaly
    Id string `json:"id"`


    // QueueId - The ID of the queue for which the anomaly was detected
    QueueId string `json:"queueId"`


    // MediaType - The media type associated with the detected anomaly
    MediaType string `json:"mediaType"`


    // DivisionId - The ID of the division in which the anomaly was detected
    DivisionId string `json:"divisionId"`


    // FalsePositive - Indicates whether the anomaly has been flagged as a false positive
    FalsePositive bool `json:"falsePositive"`


    // Metric - The name of the metric for which the anomaly was detected
    Metric string `json:"metric"`


    // DateAnomalyDetected - The date and time at which the anomaly was detected, in ISO-8601 format
    DateAnomalyDetected time.Time `json:"dateAnomalyDetected"`


    // DateIntervalStart - The beginning of the measured time period during which the anomaly was detected, in ISO-8601 format
    DateIntervalStart time.Time `json:"dateIntervalStart"`


    // DateIntervalEnd - The end of the measured time period during which the anomaly was detected, in ISO-8601 format
    DateIntervalEnd time.Time `json:"dateIntervalEnd"`


    // ConsensusAnalyses - The consensus methods used for analysis and their metadata
    ConsensusAnalyses []Anomalyconsensusanalysis `json:"consensusAnalyses"`


    // AnomalyScorePercent - Percentage value (0-100) representing confidence that the detected value is anomalous
    AnomalyScorePercent float64 `json:"anomalyScorePercent"`


    // DetectedValue - The historical value that was analyzed
    DetectedValue float64 `json:"detectedValue"`


    // LowerBound - The lower bound of the normal range
    LowerBound float64 `json:"lowerBound"`


    // UpperBound - The upper bound of the normal range
    UpperBound float64 `json:"upperBound"`


    // ExpectedValue - The expected value based on historical data
    ExpectedValue float64 `json:"expectedValue"`


    // Deviation - Percentage deviation from the expected value
    Deviation float64 `json:"deviation"`


    // DateExpires - The date and time when this anomaly expires, in ISO-8601 format
    DateExpires time.Time `json:"dateExpires"`


    // DateModified - The date and time at which this anomaly was last modified, in ISO-8601 format
    DateModified time.Time `json:"dateModified"`

}

// String returns a JSON representation of the model
func (o *Queueanomalyexport) String() string {
    
    
    
    
    
    
    
    
    
     o.ConsensusAnalyses = []Anomalyconsensusanalysis{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueanomalyexport) MarshalJSON() ([]byte, error) {
    type Alias Queueanomalyexport

    if QueueanomalyexportMarshalled {
        return []byte("{}"), nil
    }
    QueueanomalyexportMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        QueueId string `json:"queueId"`
        
        MediaType string `json:"mediaType"`
        
        DivisionId string `json:"divisionId"`
        
        FalsePositive bool `json:"falsePositive"`
        
        Metric string `json:"metric"`
        
        DateAnomalyDetected time.Time `json:"dateAnomalyDetected"`
        
        DateIntervalStart time.Time `json:"dateIntervalStart"`
        
        DateIntervalEnd time.Time `json:"dateIntervalEnd"`
        
        ConsensusAnalyses []Anomalyconsensusanalysis `json:"consensusAnalyses"`
        
        AnomalyScorePercent float64 `json:"anomalyScorePercent"`
        
        DetectedValue float64 `json:"detectedValue"`
        
        LowerBound float64 `json:"lowerBound"`
        
        UpperBound float64 `json:"upperBound"`
        
        ExpectedValue float64 `json:"expectedValue"`
        
        Deviation float64 `json:"deviation"`
        
        DateExpires time.Time `json:"dateExpires"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        ConsensusAnalyses: []Anomalyconsensusanalysis{{}},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

