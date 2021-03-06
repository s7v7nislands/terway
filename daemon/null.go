package daemon

import (
	"github.com/AliyunContainerService/terway/types"
	"fmt"
)

type NullResourceManager struct {
}

func (NullResourceManager) Allocate(context *NetworkContext, prefer string) (types.NetworkResource, error) {
	return nil, fmt.Errorf("unsupported")
}

func (NullResourceManager) Release(context *NetworkContext, resId string) error {
	return fmt.Errorf("unsupported")
}

func (NullResourceManager) GarbageCollection(inUseResList []string, expireResList []string) error {
	return fmt.Errorf("unsupported")
}