package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscheduleslotsjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscheduleslotsjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Learningscheduleslotsjobresponse
type Learningscheduleslotsjobresponse struct { 
    


    // UserIds - The user IDs to fetch the slots for.
    UserIds []string `json:"userIds"`


    // LengthInMinutes - The length in minutes of the slots.
    LengthInMinutes int `json:"lengthInMinutes"`


    // BusinessUnitId - The Business Unit Id of the users.
    BusinessUnitId string `json:"businessUnitId"`


    // ActivityCodeId - The Activity Code Id of the slots.
    ActivityCodeId string `json:"activityCodeId"`


    // SlotsType - The type of slots fetched in the job.
    SlotsType string `json:"slotsType"`


    // Results - The results of the job.
    Results []Learningscheduleslotsjobresult `json:"results"`


    

}

// String returns a JSON representation of the model
func (o *Learningscheduleslotsjobresponse) String() string {
     o.UserIds = []string{""} 
    
    
    
    
     o.Results = []Learningscheduleslotsjobresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscheduleslotsjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningscheduleslotsjobresponse

    if LearningscheduleslotsjobresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningscheduleslotsjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        UserIds []string `json:"userIds"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        BusinessUnitId string `json:"businessUnitId"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        SlotsType string `json:"slotsType"`
        
        Results []Learningscheduleslotsjobresult `json:"results"`
        *Alias
    }{

        


        
        UserIds: []string{""},
        


        


        


        


        


        
        Results: []Learningscheduleslotsjobresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

