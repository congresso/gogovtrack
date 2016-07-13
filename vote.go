package gogovtrack

// VotesResponse is
type VotesResponse struct {
	Meta  Meta    `json:"meta"`
	Votes []*Vote `json:"objects"`
}

// VoteResponse is
type VoteResponse Vote

// Vote is
type Vote struct {
	Category         string      `json:"category"`
	CategoryLabel    string      `json:"category_label"`
	Chamber          string      `json:"chamber"`
	ChamberLabel     string      `json:"chamber_label"`
	Congress         int         `json:"congress"`
	Created          string      `json:"created"`
	VoteID           int         `json:"id"`
	Link             string      `json:"link"`
	Margin           float64     `json:"margin"`
	MissingData      bool        `json:"missing_data"`
	Number           int         `json:"number"`
	Options          []Option    `json:"options"`
	PercentPlus      float64     `json:"percent_plus"`
	Question         string      `json:"question"`
	QuestionDetails  interface{} `json:"question_details"`
	RelatedAmendment interface{} `json:"related_amendment"`
	RelatedBill      interface{} `json:"related_bill"`
	Required         string      `json:"required"`
	Result           string      `json:"result"`
	Session          string      `json:"session"`
	Source           string      `json:"source"`
	SourceLabel      string      `json:"source_label"`
	TotalMinus       int         `json:"total_minus"`
	TotalOther       int         `json:"total_other"`
	TotalPlus        int         `json:"total_plus"`
	VoteType         string      `json:"vote_type"`
}

// Option is
type Option struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Vote  int    `json:"vote"`
}
