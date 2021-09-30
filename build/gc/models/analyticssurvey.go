package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticssurveyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticssurveyDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Analyticssurvey
type Analyticssurvey struct { 
    // EventTime - Specifies when an event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventTime time.Time `json:"eventTime"`


    // QueueId - The ID of the associated queue
    QueueId string `json:"queueId"`


    // SurveyCompletedDate - Completion datetime of the survey in ISO 8601 format
    SurveyCompletedDate time.Time `json:"surveyCompletedDate"`


    // SurveyFormContextId - Unique identifier for the survey form, regardless of version
    SurveyFormContextId string `json:"surveyFormContextId"`


    // SurveyFormId - ID of the survey form used
    SurveyFormId string `json:"surveyFormId"`


    // SurveyFormName - Name of the survey form used
    SurveyFormName string `json:"surveyFormName"`


    // SurveyId - ID of the survey
    SurveyId string `json:"surveyId"`


    // SurveyPromoterScore - Score of the survey used with NPS
    SurveyPromoterScore int `json:"surveyPromoterScore"`


    // SurveyStatus - The status of the survey
    SurveyStatus string `json:"surveyStatus"`


    // UserId - ID of the agent the survey was performed against
    UserId string `json:"userId"`


    // OSurveyTotalScore
    OSurveyTotalScore int `json:"oSurveyTotalScore"`

}

// String returns a JSON representation of the model
func (o *Analyticssurvey) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticssurvey) MarshalJSON() ([]byte, error) {
    type Alias Analyticssurvey

    if AnalyticssurveyMarshalled {
        return []byte("{}"), nil
    }
    AnalyticssurveyMarshalled = true

    return json.Marshal(&struct { 
        EventTime time.Time `json:"eventTime"`
        
        QueueId string `json:"queueId"`
        
        SurveyCompletedDate time.Time `json:"surveyCompletedDate"`
        
        SurveyFormContextId string `json:"surveyFormContextId"`
        
        SurveyFormId string `json:"surveyFormId"`
        
        SurveyFormName string `json:"surveyFormName"`
        
        SurveyId string `json:"surveyId"`
        
        SurveyPromoterScore int `json:"surveyPromoterScore"`
        
        SurveyStatus string `json:"surveyStatus"`
        
        UserId string `json:"userId"`
        
        OSurveyTotalScore int `json:"oSurveyTotalScore"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

