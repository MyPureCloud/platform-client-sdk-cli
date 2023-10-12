package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventDud struct { 
    


    CorrelationId string `json:"correlationId"`


    


    


    


    


    


    


    


    


    


    


    


    

}

// Event
type Event struct { 
    // Id - System-generated UUID for the event.
    Id string `json:"id"`


    


    // CustomerId - Primary identifier of the customer in the source of the events.
    CustomerId string `json:"customerId"`


    // CustomerIdType - Type of primary identifier (e.g. cookie, email, phone).
    CustomerIdType string `json:"customerIdType"`


    // Session - The session that the event belongs to.
    Session Eventsession `json:"session"`


    // EventType - The name representing the type of event.
    EventType string `json:"eventType"`


    // GenericActionEvent - Event triggered by generic actions.
    GenericActionEvent Genericactionevent `json:"genericActionEvent"`


    // OutcomeAchievedEvent - Event where a customer has achieved a specific outcome or goal.
    OutcomeAchievedEvent Outcomeachievedevent `json:"outcomeAchievedEvent"`


    // SegmentAssignedEvent - Event that represents a segment being assigned (deprecated).
    SegmentAssignedEvent Segmentassignedevent `json:"segmentAssignedEvent"`


    // SegmentAssignmentEvent - Event that represents a segment being assigned.
    SegmentAssignmentEvent Segmentassignmentevent `json:"segmentAssignmentEvent"`


    // WebActionEvent - Event triggered by web actions.
    WebActionEvent Webactionevent `json:"webActionEvent"`


    // WebEvent - Event that tracks user interactions with content in a browser such as pageviews, downloads, mobile ad clicks, etc.
    WebEvent Webevent `json:"webEvent"`


    // AppEvent - Event that tracks user interactions with content in an application such as screen views, searches, etc.
    AppEvent Appevent `json:"appEvent"`


    // CreatedDate - Timestamp indicating when the event actually took place. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Event) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Event) MarshalJSON() ([]byte, error) {
    type Alias Event

    if EventMarshalled {
        return []byte("{}"), nil
    }
    EventMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        CustomerId string `json:"customerId"`
        
        CustomerIdType string `json:"customerIdType"`
        
        Session Eventsession `json:"session"`
        
        EventType string `json:"eventType"`
        
        GenericActionEvent Genericactionevent `json:"genericActionEvent"`
        
        OutcomeAchievedEvent Outcomeachievedevent `json:"outcomeAchievedEvent"`
        
        SegmentAssignedEvent Segmentassignedevent `json:"segmentAssignedEvent"`
        
        SegmentAssignmentEvent Segmentassignmentevent `json:"segmentAssignmentEvent"`
        
        WebActionEvent Webactionevent `json:"webActionEvent"`
        
        WebEvent Webevent `json:"webEvent"`
        
        AppEvent Appevent `json:"appEvent"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

