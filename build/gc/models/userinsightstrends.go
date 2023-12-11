package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserinsightstrendsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserinsightstrendsDud struct { 
    


    

}

// Userinsightstrends
type Userinsightstrends struct { 
    // ComparativePeriod - List of trend data in the comparative period
    ComparativePeriod []Usertrenddata `json:"comparativePeriod"`


    // PrimaryPeriod - List of trend data in the primary period
    PrimaryPeriod []Usertrenddata `json:"primaryPeriod"`

}

// String returns a JSON representation of the model
func (o *Userinsightstrends) String() string {
     o.ComparativePeriod = []Usertrenddata{{}} 
     o.PrimaryPeriod = []Usertrenddata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userinsightstrends) MarshalJSON() ([]byte, error) {
    type Alias Userinsightstrends

    if UserinsightstrendsMarshalled {
        return []byte("{}"), nil
    }
    UserinsightstrendsMarshalled = true

    return json.Marshal(&struct {
        
        ComparativePeriod []Usertrenddata `json:"comparativePeriod"`
        
        PrimaryPeriod []Usertrenddata `json:"primaryPeriod"`
        *Alias
    }{

        
        ComparativePeriod: []Usertrenddata{{}},
        


        
        PrimaryPeriod: []Usertrenddata{{}},
        

        Alias: (*Alias)(u),
    })
}

