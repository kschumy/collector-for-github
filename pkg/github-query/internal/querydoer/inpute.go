package querydoer

import (
	"time"

	"github.com/collector-for-GitHub/pkg/github-query/types"
)

//type RelativeTimeProvider interface {
//	GetRelativeTime() types.RelativeTime
//}

//type QueryStringGenerator interface {
//	GenerateQueryString(time.Time) (string, error)
//}

type DoQueryInput struct {
	QueryFactory      func(time.Time) (string, error)
	RelativeTimeStart types.RelativeTime
	ObjectType        types.ObjectType
}

func DoQuery(input DoQueryInput) (*doQueryManager, error) {
	manager := doQueryManager{
		getString:        input.QueryFactory,
		currRelativeTime: input.RelativeTimeStart,
		objectType:       input.ObjectType,
	}

	err := manager.doQuery()
	if err != nil {
		return nil, err
	}
	return &manager, nil
}
