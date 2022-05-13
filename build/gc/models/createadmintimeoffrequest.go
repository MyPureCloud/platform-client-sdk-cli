package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateadmintimeoffrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateadmintimeoffrequestDud struct { 
    


    


    


    


    


    


    

}

// Createadmintimeoffrequest
type Createadmintimeoffrequest struct { 
    // Status - The status of this time off request
    Status string `json:"status"`


    // Users - A set of IDs for users to associate with this time off request
    Users []Userreference `json:"users"`


    // ActivityCodeId - The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
    ActivityCodeId string `json:"activityCodeId"`


    // Notes - Notes about the time off request
    Notes string `json:"notes"`


    // FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.
    FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`


    // PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.
    PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`


    // DailyDurationMinutes - The daily duration of this time off request in minutes
    DailyDurationMinutes int `json:"dailyDurationMinutes"`

}

// String returns a JSON representation of the model
func (o *Createadmintimeoffrequest) String() string {
    
     o.Users = []Userreference{{}} 
    
    
     o.FullDayManagementUnitDates = []string{""} 
     o.PartialDayStartDateTimes = []time.Time{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createadmintimeoffrequest) MarshalJSON() ([]byte, error) {
    type Alias Createadmintimeoffrequest

    if CreateadmintimeoffrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateadmintimeoffrequestMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        Users []Userreference `json:"users"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Notes string `json:"notes"`
        
        FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`
        
        PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`
        
        DailyDurationMinutes int `json:"dailyDurationMinutes"`
        *Alias
    }{

        


        
        Users: []Userreference{{}},
        


        


        


        
        FullDayManagementUnitDates: []string{""},
        


        
        PartialDayStartDateTimes: []time.Time{{}},
        


        

        Alias: (*Alias)(u),
    })
}

