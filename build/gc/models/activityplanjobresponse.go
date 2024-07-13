package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanjobresponse
type Activityplanjobresponse struct { 
    


    // ActivityPlan - The activity plan associated with this job
    ActivityPlan Activityplanreference `json:"activityPlan"`


    // Status - The status of the job
    Status string `json:"status"`


    // Exceptions - The list of exceptions that occurred while running this activity plan job. These are exceptions that affect individual occurrences but didn't prevent the job from completing
    Exceptions []Activityplanjobexception `json:"exceptions"`


    // VarError - Error details if status == 'Error'. These are errors that caused the job to fail to complete
    VarError Errorbody `json:"error"`


    // Occurrence - The occurrence associated with this job if type == 'DeleteOccurrence'
    Occurrence Activityplanoccurrencereference `json:"occurrence"`


    // VarType - The type of the job
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanjobresponse) String() string {
    
    
     o.Exceptions = []Activityplanjobexception{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Activityplanjobresponse

    if ActivityplanjobresponseMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        ActivityPlan Activityplanreference `json:"activityPlan"`
        
        Status string `json:"status"`
        
        Exceptions []Activityplanjobexception `json:"exceptions"`
        
        VarError Errorbody `json:"error"`
        
        Occurrence Activityplanoccurrencereference `json:"occurrence"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        


        
        Exceptions: []Activityplanjobexception{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

