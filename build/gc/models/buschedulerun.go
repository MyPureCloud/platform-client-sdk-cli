package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulerunMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulerunDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Buschedulerun
type Buschedulerun struct { 
    


    // SchedulerRunId - The scheduler run ID.  Reference this value for support
    SchedulerRunId string `json:"schedulerRunId"`


    // IntradayRescheduling - Whether this is an intraday rescheduling run
    IntradayRescheduling bool `json:"intradayRescheduling"`


    // State - The state of the generation run
    State string `json:"state"`


    // WeekCount - The number of weeks spanned by the schedule
    WeekCount int `json:"weekCount"`


    // PercentComplete - Percent completion of the schedule run
    PercentComplete float64 `json:"percentComplete"`


    // TargetWeek - The start date of the target week. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    TargetWeek time.Time `json:"targetWeek"`


    // Schedule - The generated schedule.  Null unless the schedule run is complete
    Schedule Buschedulereference `json:"schedule"`


    // ScheduleDescription - The description of the generated schedule
    ScheduleDescription string `json:"scheduleDescription"`


    // SchedulingStartTime - When the schedule generation run started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SchedulingStartTime time.Time `json:"schedulingStartTime"`


    // SchedulingStartedBy - The user who started the scheduling run
    SchedulingStartedBy Userreference `json:"schedulingStartedBy"`


    // SchedulingCanceledBy - The user who canceled the scheduling run, if applicable
    SchedulingCanceledBy Userreference `json:"schedulingCanceledBy"`


    // SchedulingCompletedTime - When the scheduling run was completed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SchedulingCompletedTime time.Time `json:"schedulingCompletedTime"`


    // MessageCount - The number of schedule generation messages for this schedule generation run
    MessageCount int `json:"messageCount"`


    // MessageSeverityCounts - The list of schedule generation message counts by severity for this schedule generation run
    MessageSeverityCounts []Schedulermessageseveritycount `json:"messageSeverityCounts"`


    // ReschedulingOptions - Rescheduling options for this run.  Null unless intradayRescheduling is true
    ReschedulingOptions Reschedulingoptionsrunresponse `json:"reschedulingOptions"`


    // ReschedulingResultExpiration - When the reschedule result will expire.  Null unless intradayRescheduling is true. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReschedulingResultExpiration time.Time `json:"reschedulingResultExpiration"`


    

}

// String returns a JSON representation of the model
func (o *Buschedulerun) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.MessageSeverityCounts = []Schedulermessageseveritycount{{}} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulerun) MarshalJSON() ([]byte, error) {
    type Alias Buschedulerun

    if BuschedulerunMarshalled {
        return []byte("{}"), nil
    }
    BuschedulerunMarshalled = true

    return json.Marshal(&struct { 
        
        
        SchedulerRunId string `json:"schedulerRunId"`
        
        IntradayRescheduling bool `json:"intradayRescheduling"`
        
        State string `json:"state"`
        
        WeekCount int `json:"weekCount"`
        
        PercentComplete float64 `json:"percentComplete"`
        
        TargetWeek time.Time `json:"targetWeek"`
        
        Schedule Buschedulereference `json:"schedule"`
        
        ScheduleDescription string `json:"scheduleDescription"`
        
        SchedulingStartTime time.Time `json:"schedulingStartTime"`
        
        SchedulingStartedBy Userreference `json:"schedulingStartedBy"`
        
        SchedulingCanceledBy Userreference `json:"schedulingCanceledBy"`
        
        SchedulingCompletedTime time.Time `json:"schedulingCompletedTime"`
        
        MessageCount int `json:"messageCount"`
        
        MessageSeverityCounts []Schedulermessageseveritycount `json:"messageSeverityCounts"`
        
        ReschedulingOptions Reschedulingoptionsrunresponse `json:"reschedulingOptions"`
        
        ReschedulingResultExpiration time.Time `json:"reschedulingResultExpiration"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        MessageSeverityCounts: []Schedulermessageseveritycount{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

