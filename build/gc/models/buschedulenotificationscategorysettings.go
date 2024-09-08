package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulenotificationscategorysettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulenotificationscategorysettingsDud struct { 
    


    


    

}

// Buschedulenotificationscategorysettings
type Buschedulenotificationscategorysettings struct { 
    // ActivityCategory - The activity category
    ActivityCategory string `json:"activityCategory"`


    // EarlyReminderEnabled - Indicates if agents should receive early schedule reminder notifications.
    EarlyReminderEnabled bool `json:"earlyReminderEnabled"`


    // OnTimeReminderEnabled - Indicates if agents should receive out of adherence notifications.
    OnTimeReminderEnabled bool `json:"onTimeReminderEnabled"`

}

// String returns a JSON representation of the model
func (o *Buschedulenotificationscategorysettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulenotificationscategorysettings) MarshalJSON() ([]byte, error) {
    type Alias Buschedulenotificationscategorysettings

    if BuschedulenotificationscategorysettingsMarshalled {
        return []byte("{}"), nil
    }
    BuschedulenotificationscategorysettingsMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCategory string `json:"activityCategory"`
        
        EarlyReminderEnabled bool `json:"earlyReminderEnabled"`
        
        OnTimeReminderEnabled bool `json:"onTimeReminderEnabled"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

