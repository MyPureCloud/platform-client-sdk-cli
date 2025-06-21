package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsevaluationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsevaluationDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Analyticsevaluation
type Analyticsevaluation struct { 
    // AssigneeApplicable - Indicates whether an assignee is applicable for the evaluation. Set to false when assignee is not applicable
    AssigneeApplicable bool `json:"assigneeApplicable"`


    // AssigneeId - UserId of the assignee
    AssigneeId string `json:"assigneeId"`


    // CalibrationId - The calibration ID used for the purpose of training evaluators
    CalibrationId string `json:"calibrationId"`


    // ContextId - A unique identifier for an evaluation form, regardless of version
    ContextId string `json:"contextId"`


    // Deleted - Whether the evaluation has been deleted
    Deleted bool `json:"deleted"`


    // EvaluationId - Unique identifier for the evaluation
    EvaluationId string `json:"evaluationId"`


    // EvaluationStatus - Status of evaluation
    EvaluationStatus string `json:"evaluationStatus"`


    // EvaluatorId - A unique identifier of the user who evaluated the interaction
    EvaluatorId string `json:"evaluatorId"`


    // EventTime - Specifies when an evaluation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventTime time.Time `json:"eventTime"`


    // FormId - ID of the evaluation form used
    FormId string `json:"formId"`


    // FormName - Name of the evaluation form used
    FormName string `json:"formName"`


    // QueueId - The ID of the associated queue
    QueueId string `json:"queueId"`


    // Released - Whether the evaluation has been released
    Released bool `json:"released"`


    // Rescored - Whether the evaluation has been rescored at least once
    Rescored bool `json:"rescored"`


    // SystemSubmitted - Whether the evaluation was auto submitted by the system
    SystemSubmitted bool `json:"systemSubmitted"`


    // UserId - ID of the agent the evaluation was performed against
    UserId string `json:"userId"`


    // OTotalCriticalScore
    OTotalCriticalScore int `json:"oTotalCriticalScore"`


    // OTotalScore
    OTotalScore int `json:"oTotalScore"`

}

// String returns a JSON representation of the model
func (o *Analyticsevaluation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsevaluation) MarshalJSON() ([]byte, error) {
    type Alias Analyticsevaluation

    if AnalyticsevaluationMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsevaluationMarshalled = true

    return json.Marshal(&struct {
        
        AssigneeApplicable bool `json:"assigneeApplicable"`
        
        AssigneeId string `json:"assigneeId"`
        
        CalibrationId string `json:"calibrationId"`
        
        ContextId string `json:"contextId"`
        
        Deleted bool `json:"deleted"`
        
        EvaluationId string `json:"evaluationId"`
        
        EvaluationStatus string `json:"evaluationStatus"`
        
        EvaluatorId string `json:"evaluatorId"`
        
        EventTime time.Time `json:"eventTime"`
        
        FormId string `json:"formId"`
        
        FormName string `json:"formName"`
        
        QueueId string `json:"queueId"`
        
        Released bool `json:"released"`
        
        Rescored bool `json:"rescored"`
        
        SystemSubmitted bool `json:"systemSubmitted"`
        
        UserId string `json:"userId"`
        
        OTotalCriticalScore int `json:"oTotalCriticalScore"`
        
        OTotalScore int `json:"oTotalScore"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

