package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulenotificationssettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulenotificationssettingsresponseDud struct { 
    


    

}

// Buschedulenotificationssettingsresponse
type Buschedulenotificationssettingsresponse struct { 
    // EarlyReminderMinutes - The number of minutes prior to the scheduled event to display an early reminder notification
    EarlyReminderMinutes int `json:"earlyReminderMinutes"`


    // ActivityCategorySettings - List of activity category notification settings
    ActivityCategorySettings []Buschedulenotificationscategorysettings `json:"activityCategorySettings"`

}

// String returns a JSON representation of the model
func (o *Buschedulenotificationssettingsresponse) String() string {
    
     o.ActivityCategorySettings = []Buschedulenotificationscategorysettings{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulenotificationssettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Buschedulenotificationssettingsresponse

    if BuschedulenotificationssettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    BuschedulenotificationssettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        EarlyReminderMinutes int `json:"earlyReminderMinutes"`
        
        ActivityCategorySettings []Buschedulenotificationscategorysettings `json:"activityCategorySettings"`
        *Alias
    }{

        


        
        ActivityCategorySettings: []Buschedulenotificationscategorysettings{{}},
        

        Alias: (*Alias)(u),
    })
}

