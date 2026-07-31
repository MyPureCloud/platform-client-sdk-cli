package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuconverttimeofflimitgranularityjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuconverttimeofflimitgranularityjobrequestDud struct { 
    


    

}

// Buconverttimeofflimitgranularityjobrequest
type Buconverttimeofflimitgranularityjobrequest struct { 
    // Granularity - Granularity to convert the time-off limit to
    Granularity string `json:"granularity"`


    // FullDayTimeOffStartTime - The start time of full day time-off requests associated with this limit interval in HH:mm format.The value can be set only once when converting the time-off limit from daily granularity to fifteen minutes.Setting this value is allowed only for time-off limit with fifteen minutes granularity.When converting time-off limit from fifteen minutes to daily granularity, the existing value is reset.
    FullDayTimeOffStartTime string `json:"fullDayTimeOffStartTime"`

}

// String returns a JSON representation of the model
func (o *Buconverttimeofflimitgranularityjobrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buconverttimeofflimitgranularityjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Buconverttimeofflimitgranularityjobrequest

    if BuconverttimeofflimitgranularityjobrequestMarshalled {
        return []byte("{}"), nil
    }
    BuconverttimeofflimitgranularityjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Granularity string `json:"granularity"`
        
        FullDayTimeOffStartTime string `json:"fullDayTimeOffStartTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

