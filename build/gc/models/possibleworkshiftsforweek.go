package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PossibleworkshiftsforweekMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PossibleworkshiftsforweekDud struct { 
    


    

}

// Possibleworkshiftsforweek
type Possibleworkshiftsforweek struct { 
    // Id - ID of this possible weekly shift
    Id int `json:"id"`


    // DailyPossibleShifts - Daily shifts in this possible weekly shift
    DailyPossibleShifts []Dailypossibleshift `json:"dailyPossibleShifts"`

}

// String returns a JSON representation of the model
func (o *Possibleworkshiftsforweek) String() string {
    
     o.DailyPossibleShifts = []Dailypossibleshift{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Possibleworkshiftsforweek) MarshalJSON() ([]byte, error) {
    type Alias Possibleworkshiftsforweek

    if PossibleworkshiftsforweekMarshalled {
        return []byte("{}"), nil
    }
    PossibleworkshiftsforweekMarshalled = true

    return json.Marshal(&struct {
        
        Id int `json:"id"`
        
        DailyPossibleShifts []Dailypossibleshift `json:"dailyPossibleShifts"`
        *Alias
    }{

        


        
        DailyPossibleShifts: []Dailypossibleshift{{}},
        

        Alias: (*Alias)(u),
    })
}

