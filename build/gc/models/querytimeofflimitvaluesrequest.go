package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerytimeofflimitvaluesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerytimeofflimitvaluesrequestDud struct { 
    


    


    

}

// Querytimeofflimitvaluesrequest - This sets up a filter to request date ranges of time off limit, allocated and waitlisted minutes
type Querytimeofflimitvaluesrequest struct { 
    // TimeOffLimitId - The time off limit object id to retrieve values for. Required if activityCodeId is not specified
    TimeOffLimitId string `json:"timeOffLimitId"`


    // ActivityCodeId - The activity code id to filter the affected limit objects by. Required if timeOffLimitId is not specified
    ActivityCodeId string `json:"activityCodeId"`


    // DateRanges - The list of the date ranges to return time off limit, allocated and waitlisted minutes.
    DateRanges []Localdaterange `json:"dateRanges"`

}

// String returns a JSON representation of the model
func (o *Querytimeofflimitvaluesrequest) String() string {
    
    
    
    
    
    
    
    
    
    
     o.DateRanges = []Localdaterange{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querytimeofflimitvaluesrequest) MarshalJSON() ([]byte, error) {
    type Alias Querytimeofflimitvaluesrequest

    if QuerytimeofflimitvaluesrequestMarshalled {
        return []byte("{}"), nil
    }
    QuerytimeofflimitvaluesrequestMarshalled = true

    return json.Marshal(&struct { 
        TimeOffLimitId string `json:"timeOffLimitId"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        DateRanges []Localdaterange `json:"dateRanges"`
        
        *Alias
    }{
        

        

        

        

        

        
        DateRanges: []Localdaterange{{}},
        

        
        Alias: (*Alias)(u),
    })
}

