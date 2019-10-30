package builder

import (
	"github.com/tiezhong2004/go-kairosdb/builder/grouper"
)

type GroupByType string

const (
	GROUP_TAGS  GroupByType = "tag"
	GROUP_TIME  GroupByType = "time"
	GROUP_VALUE  GroupByType = "value"
	GROUP_BIN  GroupByType = "bin"
)


func CreateTagsGroupBy(tags []string) Grouper {
	return grouper.NewTagsGroup(tags)
}
