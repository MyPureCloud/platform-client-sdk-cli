package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanoccurrencesdeletionjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanoccurrencesdeletionjobresponseDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanoccurrencesdeletionjobresponse
type Activityplanoccurrencesdeletionjobresponse struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Status - The status of the job
    Status string `json:"status"`


    // Exceptions - The list of exceptions that occurred while running this activity plan job. These are exceptions that affect individual occurrences but didn't prevent the job from completing
    Exceptions []Activityplanjobexception `json:"exceptions"`


    // VarError - Error details if status == 'Error'. These are errors that caused the job to fail to complete
    VarError Errorbody `json:"error"`


    // ActivityPlan - The activity plan associated with this job
    ActivityPlan Activityplanstructurewithoccurrencesreference `json:"activityPlan"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanoccurrencesdeletionjobresponse) String() string {
    
    
     o.Exceptions = []Activityplanjobexception{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanoccurrencesdeletionjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Activityplanoccurrencesdeletionjobresponse

    if ActivityplanoccurrencesdeletionjobresponseMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanoccurrencesdeletionjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        Exceptions []Activityplanjobexception `json:"exceptions"`
        
        VarError Errorbody `json:"error"`
        
        ActivityPlan Activityplanstructurewithoccurrencesreference `json:"activityPlan"`
        *Alias
    }{

        


        


        
        Exceptions: []Activityplanjobexception{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

