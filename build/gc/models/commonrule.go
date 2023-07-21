package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommonruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommonruleDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Commonrule
type Commonrule struct { 
    


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


    // VarType - The type of the rule.
    VarType string `json:"type"`


    // InAlarm - Indicates if the rule is in alarm state.
    InAlarm bool `json:"inAlarm"`


    // User - The entity that created the rule.
    User Userreference `json:"user"`


    // Version - The current version number of the rule.
    Version int `json:"version"`


    // DateCreated - The creation date of the rule when the rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateLastModified - The timestamp of the last update to the rule. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateLastModified time.Time `json:"dateLastModified"`


    

}

// String returns a JSON representation of the model
func (o *Commonrule) String() string {
    
    
    
     o.Notifications = []Alertnotification{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commonrule) MarshalJSON() ([]byte, error) {
    type Alias Commonrule

    if CommonruleMarshalled {
        return []byte("{}"), nil
    }
    CommonruleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Enabled bool `json:"enabled"`
        
        Notifications []Alertnotification `json:"notifications"`
        
        SendExitingAlarmNotifications bool `json:"sendExitingAlarmNotifications"`
        
        WaitBetweenNotificationMs int `json:"waitBetweenNotificationMs"`
        
        Conditions Commonruleconditions `json:"conditions"`
        
        VarType string `json:"type"`
        
        InAlarm bool `json:"inAlarm"`
        
        User Userreference `json:"user"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateLastModified time.Time `json:"dateLastModified"`
        *Alias
    }{

        


        


        


        


        
        Notifications: []Alertnotification{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

