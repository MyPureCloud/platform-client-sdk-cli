package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommonalertMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommonalertDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Commonalert
type Commonalert struct { 
    


    // Name
    Name string `json:"name"`


    // User - The user who created the rule that triggered the alert.
    User Userreference `json:"user"`


    // Rule - The properties of the rule that triggered the alert.
    Rule Alertruleproperties `json:"rule"`


    // Notifications - The collection of notification methods and the ids of users who were notified by those methods.
    Notifications []Alertnotification `json:"notifications"`


    // DateStart - The timestamp of when the alert was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The timestamp of when the alert ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // Active - Indicates if an alert is currently active.
    Active bool `json:"active"`


    // Unread - Indicates if an alert has not been read.
    Unread bool `json:"unread"`


    // WaitBetweenNotificationMs - The amount of time to wait between notification. Time is in milliseconds.
    WaitBetweenNotificationMs int `json:"waitBetweenNotificationMs"`


    // Muted - Flag indicating if the alert is in a muted state.
    Muted bool `json:"muted"`


    // Snoozed - Flag indicating if the alert is in a snoozed state.
    Snoozed bool `json:"snoozed"`


    // DateMutedUntil - Timestamp of when the mute status of the alert should end. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateMutedUntil time.Time `json:"dateMutedUntil"`


    // DateSnoozedUntil - Timestamp of when the snooze status of the alert should end. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateSnoozedUntil time.Time `json:"dateSnoozedUntil"`


    // Conditions - The conditions that make up the rule.
    Conditions Commonruleconditions `json:"conditions"`


    // ConversationId - The id of the conversation instance that caused the alert to trigger.
    ConversationId string `json:"conversationId"`


    // RuleUri
    RuleUri string `json:"ruleUri"`


    

}

// String returns a JSON representation of the model
func (o *Commonalert) String() string {
    
    
    
     o.Notifications = []Alertnotification{{}} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commonalert) MarshalJSON() ([]byte, error) {
    type Alias Commonalert

    if CommonalertMarshalled {
        return []byte("{}"), nil
    }
    CommonalertMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        User Userreference `json:"user"`
        
        Rule Alertruleproperties `json:"rule"`
        
        Notifications []Alertnotification `json:"notifications"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Active bool `json:"active"`
        
        Unread bool `json:"unread"`
        
        WaitBetweenNotificationMs int `json:"waitBetweenNotificationMs"`
        
        Muted bool `json:"muted"`
        
        Snoozed bool `json:"snoozed"`
        
        DateMutedUntil time.Time `json:"dateMutedUntil"`
        
        DateSnoozedUntil time.Time `json:"dateSnoozedUntil"`
        
        Conditions Commonruleconditions `json:"conditions"`
        
        ConversationId string `json:"conversationId"`
        
        RuleUri string `json:"ruleUri"`
        *Alias
    }{

        


        


        


        


        
        Notifications: []Alertnotification{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

