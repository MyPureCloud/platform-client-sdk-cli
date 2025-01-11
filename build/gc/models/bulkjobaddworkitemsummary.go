package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobaddworkitemsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobaddworkitemsummaryDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Bulkjobaddworkitemsummary
type Bulkjobaddworkitemsummary struct { 
    


    // Name
    Name string `json:"name"`


    // Workbin - The workbin of the workitem.
    Workbin Workbinreference `json:"workbin"`


    // VarType - The worktype of the workitem.
    VarType Worktypereference `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Bulkjobaddworkitemsummary) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobaddworkitemsummary) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobaddworkitemsummary

    if BulkjobaddworkitemsummaryMarshalled {
        return []byte("{}"), nil
    }
    BulkjobaddworkitemsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Workbin Workbinreference `json:"workbin"`
        
        VarType Worktypereference `json:"type"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

