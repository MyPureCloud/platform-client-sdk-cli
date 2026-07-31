package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuconverttimeofflimitgranularityjobprogressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuconverttimeofflimitgranularityjobprogressDud struct { 
    


    


    


    

}

// Buconverttimeofflimitgranularityjobprogress
type Buconverttimeofflimitgranularityjobprogress struct { 
    // DateEarliestComplete - Earliest date completed for time-off limit granularity conversion. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEarliestComplete time.Time `json:"dateEarliestComplete"`


    // DateLatestComplete - Latest date completed for time-off limit granularity conversion. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateLatestComplete time.Time `json:"dateLatestComplete"`


    // NumberOfDaysComplete - Number of days completed for time-off limit granularity conversion
    NumberOfDaysComplete int `json:"numberOfDaysComplete"`


    // PercentageComplete - Percentage completed for time-off limit granularity conversion
    PercentageComplete int `json:"percentageComplete"`

}

// String returns a JSON representation of the model
func (o *Buconverttimeofflimitgranularityjobprogress) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buconverttimeofflimitgranularityjobprogress) MarshalJSON() ([]byte, error) {
    type Alias Buconverttimeofflimitgranularityjobprogress

    if BuconverttimeofflimitgranularityjobprogressMarshalled {
        return []byte("{}"), nil
    }
    BuconverttimeofflimitgranularityjobprogressMarshalled = true

    return json.Marshal(&struct {
        
        DateEarliestComplete time.Time `json:"dateEarliestComplete"`
        
        DateLatestComplete time.Time `json:"dateLatestComplete"`
        
        NumberOfDaysComplete int `json:"numberOfDaysComplete"`
        
        PercentageComplete int `json:"percentageComplete"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

