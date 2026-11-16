package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ModifiablerulepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ModifiablerulepropertiesDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Modifiableruleproperties
type Modifiableruleproperties struct { 
    


    // Name - Name of the rule
    Name string `json:"name"`


    // Description - The description of the rule.
    Description string `json:"description"`


    // Enabled - Indicates if the rule is enabled.
    Enabled bool `json:"enabled"`


    // Notifications - The alert notification types to trigger when alarm state changes as well as the users they will be sent to.
    Notifications []Alertnotification `json:"notifications"`


    // SendExitingAlarmNotifications - Indicates if the alert will send a notification when it is closed.
    SendExitingAlarmNotifications bool `json:"sendExitingAlarmNotifications"`


    // WaitBetweenNotificationMs - The amount of time in milliseconds to wait between notification.
    WaitBetweenNotificationMs int `json:"waitBetweenNotificationMs"`


    // Conditions - The set of metric conditions that would trigger an alert.
    Conditions Commonruleconditions `json:"conditions"`


    // Muted - Flag that indicates if the rule is muted.
    Muted bool `json:"muted"`


    // DateStartMute - Timestamp of when the mute status of the rule should start. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStartMute time.Time `json:"dateStartMute"`


    // DateEndMute - Timestamp of when the mute status of the rule should end. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEndMute time.Time `json:"dateEndMute"`


    

}

// String returns a JSON representation of the model
func (o *Modifiableruleproperties) String() string {
    
    
    
     o.Notifications = []Alertnotification{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Modifiableruleproperties) MarshalJSON() ([]byte, error) {
    type Alias Modifiableruleproperties

    if ModifiablerulepropertiesMarshalled {
        return []byte("{}"), nil
    }
    ModifiablerulepropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Enabled bool `json:"enabled"`
        
        Notifications []Alertnotification `json:"notifications"`
        
        SendExitingAlarmNotifications bool `json:"sendExitingAlarmNotifications"`
        
        WaitBetweenNotificationMs int `json:"waitBetweenNotificationMs"`
        
        Conditions Commonruleconditions `json:"conditions"`
        
        Muted bool `json:"muted"`
        
        DateStartMute time.Time `json:"dateStartMute"`
        
        DateEndMute time.Time `json:"dateEndMute"`
        *Alias
    }{

        


        


        


        


        
        Notifications: []Alertnotification{{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

