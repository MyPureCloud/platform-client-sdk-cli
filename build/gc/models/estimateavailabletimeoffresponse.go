package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EstimateavailabletimeoffresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EstimateavailabletimeoffresponseDud struct { 
    


    


    


    


    

}

// Estimateavailabletimeoffresponse
type Estimateavailabletimeoffresponse struct { 
    // FullDayDates - Full day dates. partialDayDates must be empty if this field is populated
    FullDayDates []Estimateavailablefulldaytimeoffresponse `json:"fullDayDates"`


    // PartialDayDates - Partial day dates. fullDayDates must be empty if this field is populated
    PartialDayDates []Estimateavailablepartialdaytimeoffresponse `json:"partialDayDates"`


    // User - The user to whom the time off request belongs
    User Userreference `json:"user"`


    // ActivityCodeId - The ID of the activity code associated with the time off request. Activity code must be of the TimeOff category
    ActivityCodeId string `json:"activityCodeId"`


    // Paid - Whether this estimate is for a paid time off request
    Paid bool `json:"paid"`

}

// String returns a JSON representation of the model
func (o *Estimateavailabletimeoffresponse) String() string {
     o.FullDayDates = []Estimateavailablefulldaytimeoffresponse{{}} 
     o.PartialDayDates = []Estimateavailablepartialdaytimeoffresponse{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Estimateavailabletimeoffresponse) MarshalJSON() ([]byte, error) {
    type Alias Estimateavailabletimeoffresponse

    if EstimateavailabletimeoffresponseMarshalled {
        return []byte("{}"), nil
    }
    EstimateavailabletimeoffresponseMarshalled = true

    return json.Marshal(&struct {
        
        FullDayDates []Estimateavailablefulldaytimeoffresponse `json:"fullDayDates"`
        
        PartialDayDates []Estimateavailablepartialdaytimeoffresponse `json:"partialDayDates"`
        
        User Userreference `json:"user"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        *Alias
    }{

        
        FullDayDates: []Estimateavailablefulldaytimeoffresponse{{}},
        


        
        PartialDayDates: []Estimateavailablepartialdaytimeoffresponse{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

