package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignruledatetimeconditionparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignruledatetimeconditionparametersDud struct { 
    


    


    


    


    


    

}

// Campaignruledatetimeconditionparameters
type Campaignruledatetimeconditionparameters struct { 
    // Inverted - If true, inverts the result of evaluating this sub-condition. Default is false
    Inverted bool `json:"inverted"`


    // TimeOfDay - Parameters for \"timeOfDay\" conditionType
    TimeOfDay Campaignruletimeofdayparameters `json:"timeOfDay"`


    // DayOfWeek - Parameters for \"dayOfWeek\" conditionType
    DayOfWeek Campaignruledayofweekparameters `json:"dayOfWeek"`


    // DayOfMonth - Parameters for \"dayOfMonth\" conditionType
    DayOfMonth Campaignruledayofmonthparameters `json:"dayOfMonth"`


    // SpecificDate - Parameters for \"specificDate\" conditionType
    SpecificDate Campaignrulespecificdateparameters `json:"specificDate"`


    // WeekDayOfMonth - Parameters for \"weekDayOfMonth\" conditionType
    WeekDayOfMonth Campaignruleweekdayofmonthparameters `json:"weekDayOfMonth"`

}

// String returns a JSON representation of the model
func (o *Campaignruledatetimeconditionparameters) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignruledatetimeconditionparameters) MarshalJSON() ([]byte, error) {
    type Alias Campaignruledatetimeconditionparameters

    if CampaignruledatetimeconditionparametersMarshalled {
        return []byte("{}"), nil
    }
    CampaignruledatetimeconditionparametersMarshalled = true

    return json.Marshal(&struct {
        
        Inverted bool `json:"inverted"`
        
        TimeOfDay Campaignruletimeofdayparameters `json:"timeOfDay"`
        
        DayOfWeek Campaignruledayofweekparameters `json:"dayOfWeek"`
        
        DayOfMonth Campaignruledayofmonthparameters `json:"dayOfMonth"`
        
        SpecificDate Campaignrulespecificdateparameters `json:"specificDate"`
        
        WeekDayOfMonth Campaignruleweekdayofmonthparameters `json:"weekDayOfMonth"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

