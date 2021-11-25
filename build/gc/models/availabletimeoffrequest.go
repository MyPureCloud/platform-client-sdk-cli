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


    // DateRanges - A list of date ranges of available time off minutes.
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

