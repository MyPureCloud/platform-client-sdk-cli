package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EstimateavailabletimeoffrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EstimateavailabletimeoffrequestDud struct { 
    


    


    


    

}

// Estimateavailabletimeoffrequest
type Estimateavailabletimeoffrequest struct { 
    // FullDayDates - Full day dates. partialDayDates must be empty if this field is populated
    FullDayDates []Estimateavailablefulldaytimeoffrequest `json:"fullDayDates"`


    // PartialDayDates - Partial day dates. fullDayDates must be empty if this field is populated
    PartialDayDates []Estimateavailablepartialdaytimeoffrequest `json:"partialDayDates"`


    // ActivityCodeId - The ID of the activity code associated with the time off request. Activity code must be of the TimeOff category
    ActivityCodeId string `json:"activityCodeId"`


    // Paid - Whether this estimate is for a paid time off request
    Paid bool `json:"paid"`

}

// String returns a JSON representation of the model
func (o *Estimateavailabletimeoffrequest) String() string {
     o.FullDayDates = []Estimateavailablefulldaytimeoffrequest{{}} 
     o.PartialDayDates = []Estimateavailablepartialdaytimeoffrequest{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Estimateavailabletimeoffrequest) MarshalJSON() ([]byte, error) {
    type Alias Estimateavailabletimeoffrequest

    if EstimateavailabletimeoffrequestMarshalled {
        return []byte("{}"), nil
    }
    EstimateavailabletimeoffrequestMarshalled = true

    return json.Marshal(&struct {
        
        FullDayDates []Estimateavailablefulldaytimeoffrequest `json:"fullDayDates"`
        
        PartialDayDates []Estimateavailablepartialdaytimeoffrequest `json:"partialDayDates"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        *Alias
    }{

        
        FullDayDates: []Estimateavailablefulldaytimeoffrequest{{}},
        


        
        PartialDayDates: []Estimateavailablepartialdaytimeoffrequest{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

