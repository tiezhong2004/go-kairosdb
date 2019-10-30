package grouper

import "github.com/tiezhong2004/go-kairosdb/builder/utils"

type range_size2	 struct {
	Value int            `json:"value,omitempty"`
	Unit  utils.TimeUnit `json:"unit,omitempty"`
}

type std_Grouper struct {
	GPName             string        `json:"name,omitempty"`
	Tags             []string        `json:"tags,omitempty"`
	Group_count      int64           `json:"group_count,omitempty"`
	Range_size1      int64           `json:"range_size,omitempty"`
	Range_size2      range_size2     `json:"range_size,omitempty"`
}


func NewTagsGroup(tags []string) *std_Grouper {
	return &std_Grouper{
		GPName: "tag",
		Tags: tags,
	}
}


func (gp *std_Grouper) Name() string {
	return gp.GPName
}

func (gp *std_Grouper) Validate() error {
	if gp.Tags == nil {
		return ErrorGroupByNameInvalid
	}

	return nil
}
