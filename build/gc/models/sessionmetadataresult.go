package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionmetadataresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionmetadataresultDud struct { 
    


    


    


    


    


    

}

// Sessionmetadataresult
type Sessionmetadataresult struct { 
    // SessionInfo - Information about the continuous forecast session
    SessionInfo Sessioninfo `json:"sessionInfo"`


    // Snapshots - Captured snapshots
    Snapshots []Snapshots `json:"snapshots"`


    // DateForecastStart - Start date of the forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateForecastStart time.Time `json:"dateForecastStart"`


    // DateHistoricalStart - Start date of the oldest available historical week. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateHistoricalStart time.Time `json:"dateHistoricalStart"`


    // AggregateOfferedHistoricalAvailability - Total historical availability for offered metric across all planning groups
    AggregateOfferedHistoricalAvailability Aggregatehistoricalavailability `json:"aggregateOfferedHistoricalAvailability"`


    // AggregateAverageHandleTimeHistoricalAvailability - Total historical availability for average handle time metric across all planning groups
    AggregateAverageHandleTimeHistoricalAvailability Aggregatehistoricalavailability `json:"aggregateAverageHandleTimeHistoricalAvailability"`

}

// String returns a JSON representation of the model
func (o *Sessionmetadataresult) String() string {
    
     o.Snapshots = []Snapshots{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionmetadataresult) MarshalJSON() ([]byte, error) {
    type Alias Sessionmetadataresult

    if SessionmetadataresultMarshalled {
        return []byte("{}"), nil
    }
    SessionmetadataresultMarshalled = true

    return json.Marshal(&struct {
        
        SessionInfo Sessioninfo `json:"sessionInfo"`
        
        Snapshots []Snapshots `json:"snapshots"`
        
        DateForecastStart time.Time `json:"dateForecastStart"`
        
        DateHistoricalStart time.Time `json:"dateHistoricalStart"`
        
        AggregateOfferedHistoricalAvailability Aggregatehistoricalavailability `json:"aggregateOfferedHistoricalAvailability"`
        
        AggregateAverageHandleTimeHistoricalAvailability Aggregatehistoricalavailability `json:"aggregateAverageHandleTimeHistoricalAvailability"`
        *Alias
    }{

        


        
        Snapshots: []Snapshots{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

