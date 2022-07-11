package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailabletimeoffrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailabletimeoffrequestDud struct { 
    


    

}

// Availabletimeoffrequest
type Availabletimeoffrequest struct { 
    // ActivityCodeId - The ID for activity code to query available time off minutes
    ActivityCodeId string `json:"activityCodeId"`


    // DateRanges - A list of date ranges of available time off minutes. A maximum number of date ranges is 30. The maximum total number of days in all ranges is 366. If no ranges are specified, then only the presence of the associated time off limit object will be checked. In such case, if the association exists, then the response will contain a list with of a single element filled with timeOffLimitId only.
    DateRanges []Localdaterange `json:"dateRanges"`

}

// String returns a JSON representation of the model
func (o *Availabletimeoffrequest) String() string {
    
     o.DateRanges = []Localdaterange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availabletimeoffrequest) MarshalJSON() ([]byte, error) {
    type Alias Availabletimeoffrequest

    if AvailabletimeoffrequestMarshalled {
        return []byte("{}"), nil
    }
    AvailabletimeoffrequestMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCodeId string `json:"activityCodeId"`
        
        DateRanges []Localdaterange `json:"dateRanges"`
        *Alias
    }{

        


        
        DateRanges: []Localdaterange{{}},
        

        Alias: (*Alias)(u),
    })
}

