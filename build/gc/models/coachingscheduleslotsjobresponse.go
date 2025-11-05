package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingscheduleslotsjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingscheduleslotsjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Coachingscheduleslotsjobresponse
type Coachingscheduleslotsjobresponse struct { 
    


    // AttendeeIds - The attendee IDs to fetch the slots for.
    AttendeeIds []string `json:"attendeeIds"`


    // FacilitatorIds - The facilitator IDs to fetch the slots for.
    FacilitatorIds []string `json:"facilitatorIds"`


    // LengthInMinutes - The length in minutes of the slots.
    LengthInMinutes int `json:"lengthInMinutes"`


    // BusinessUnitId - The Business Unit Id of the users.
    BusinessUnitId string `json:"businessUnitId"`


    // ActivityCodeId - The Activity Code Id of the slots.
    ActivityCodeId string `json:"activityCodeId"`


    // Results - The results of the job.
    Results []Coachingscheduleslotsjobresult `json:"results"`


    // SlotsType - The type of slots fetched in the job.
    SlotsType string `json:"slotsType"`


    

}

// String returns a JSON representation of the model
func (o *Coachingscheduleslotsjobresponse) String() string {
     o.AttendeeIds = []string{""} 
     o.FacilitatorIds = []string{""} 
    
    
    
     o.Results = []Coachingscheduleslotsjobresult{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingscheduleslotsjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Coachingscheduleslotsjobresponse

    if CoachingscheduleslotsjobresponseMarshalled {
        return []byte("{}"), nil
    }
    CoachingscheduleslotsjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        AttendeeIds []string `json:"attendeeIds"`
        
        FacilitatorIds []string `json:"facilitatorIds"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        BusinessUnitId string `json:"businessUnitId"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Results []Coachingscheduleslotsjobresult `json:"results"`
        
        SlotsType string `json:"slotsType"`
        *Alias
    }{

        


        
        AttendeeIds: []string{""},
        


        
        FacilitatorIds: []string{""},
        


        


        


        


        
        Results: []Coachingscheduleslotsjobresult{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

