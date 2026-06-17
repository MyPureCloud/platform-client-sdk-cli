package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidgroupresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidgroupresponseDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Schedulebidgroupresponse
type Schedulebidgroupresponse struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // ScheduleBidGroup - The schedule bid group
    ScheduleBidGroup Schedulebidgroup `json:"scheduleBidGroup"`


    // Metadata - The metadata of the bid group
    Metadata Workplanbidmetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Schedulebidgroupresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidgroupresponse) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidgroupresponse

    if SchedulebidgroupresponseMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidgroupresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ScheduleBidGroup Schedulebidgroup `json:"scheduleBidGroup"`
        
        Metadata Workplanbidmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

