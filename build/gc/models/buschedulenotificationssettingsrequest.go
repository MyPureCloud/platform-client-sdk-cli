package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulenotificationssettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulenotificationssettingsrequestDud struct { 
    


    

}

// Buschedulenotificationssettingsrequest
type Buschedulenotificationssettingsrequest struct { 
    // EarlyReminderMinutes - The number of minutes prior to the scheduled event to display an early reminder notification
    EarlyReminderMinutes int `json:"earlyReminderMinutes"`


    // ActivityCategorySettings - List of activity category notification settings
    ActivityCategorySettings []Buschedulenotificationscategorysettings `json:"activityCategorySettings"`

}

// String returns a JSON representation of the model
func (o *Buschedulenotificationssettingsrequest) String() string {
    
     o.ActivityCategorySettings = []Buschedulenotificationscategorysettings{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulenotificationssettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Buschedulenotificationssettingsrequest

    if BuschedulenotificationssettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    BuschedulenotificationssettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        EarlyReminderMinutes int `json:"earlyReminderMinutes"`
        
        ActivityCategorySettings []Buschedulenotificationscategorysettings `json:"activityCategorySettings"`
        *Alias
    }{

        


        
        ActivityCategorySettings: []Buschedulenotificationscategorysettings{{}},
        

        Alias: (*Alias)(u),
    })
}

