package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetrulesqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetrulesqueryDud struct { 
    


    


    


    


    


    


    


    


    

}

// Getrulesquery
type Getrulesquery struct { 
    // RuleType - The rule type of the alerts the query will return
    RuleType string `json:"ruleType"`


    // QueryType - The type of query being performed.
    QueryType string `json:"queryType"`


    // EnabledType - The state of the rule the query will return.  The accepted choices are Enabled, Disabled, or All
    EnabledType string `json:"enabledType"`


    // PageNumber - The page number of the queried response
    PageNumber int `json:"pageNumber"`


    // PageSize - The number of entities to return of the queried response.  The max is 25
    PageSize int `json:"pageSize"`


    // SortBy - The field to sort responses by.  The accepted choices are Name and DateStart
    SortBy string `json:"sortBy"`


    // SortOrder - The order in which response will be sorted.  The accepted choices are Asc and Desc
    SortOrder string `json:"sortOrder"`


    // RuleName - The name of the rule being queries.
    RuleName string `json:"ruleName"`


    // NameSearchType - Specifies how strict the name search needs to be. Expected values are Exact and Contains if querying by name.
    NameSearchType string `json:"nameSearchType"`

}

// String returns a JSON representation of the model
func (o *Getrulesquery) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getrulesquery) MarshalJSON() ([]byte, error) {
    type Alias Getrulesquery

    if GetrulesqueryMarshalled {
        return []byte("{}"), nil
    }
    GetrulesqueryMarshalled = true

    return json.Marshal(&struct {
        
        RuleType string `json:"ruleType"`
        
        QueryType string `json:"queryType"`
        
        EnabledType string `json:"enabledType"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        SortBy string `json:"sortBy"`
        
        SortOrder string `json:"sortOrder"`
        
        RuleName string `json:"ruleName"`
        
        NameSearchType string `json:"nameSearchType"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

