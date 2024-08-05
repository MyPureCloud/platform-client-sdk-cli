package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidgroupresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidgroupresponseDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Workplanbidgroupresponse
type Workplanbidgroupresponse struct { 
    


    // WorkPlanBidGroup - The work plan bid group
    WorkPlanBidGroup Workplanbidgroup `json:"workPlanBidGroup"`


    // Metadata - The meta data of the bid group
    Metadata Workplanbidmetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Workplanbidgroupresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidgroupresponse) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidgroupresponse

    if WorkplanbidgroupresponseMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidgroupresponseMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlanBidGroup Workplanbidgroup `json:"workPlanBidGroup"`
        
        Metadata Workplanbidmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

