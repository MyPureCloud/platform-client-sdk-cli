package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscheduleslotsqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscheduleslotsqueryresponseDud struct { 
    


    

}

// Learningscheduleslotsqueryresponse
type Learningscheduleslotsqueryresponse struct { 
    // SuggestedSlots - List of slots where Learning activity can be scheduled
    SuggestedSlots []Learningslot `json:"suggestedSlots"`


    // WfmScheduleActivities - Detailed data for WFM scheduled activities
    WfmScheduleActivities []Learningslotwfmscheduleactivity `json:"wfmScheduleActivities"`

}

// String returns a JSON representation of the model
func (o *Learningscheduleslotsqueryresponse) String() string {
     o.SuggestedSlots = []Learningslot{{}} 
     o.WfmScheduleActivities = []Learningslotwfmscheduleactivity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscheduleslotsqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningscheduleslotsqueryresponse

    if LearningscheduleslotsqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningscheduleslotsqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        SuggestedSlots []Learningslot `json:"suggestedSlots"`
        
        WfmScheduleActivities []Learningslotwfmscheduleactivity `json:"wfmScheduleActivities"`
        *Alias
    }{

        
        SuggestedSlots: []Learningslot{{}},
        


        
        WfmScheduleActivities: []Learningslotwfmscheduleactivity{{}},
        

        Alias: (*Alias)(u),
    })
}

