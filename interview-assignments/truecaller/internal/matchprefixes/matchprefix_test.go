package matchprefixes

import (
	"testing"

	"dkgosql.com/match-prefix-service/mocks"
	"github.com/golang/mock/gomock"
)

func TestFindPrefix(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	prefixFinder := mocks.NewMockPrefixFinder(mockCtrl)
	
	prefixFinder.EXPECT().FindPrefix("xyz").Return(nil).MinTimes(1)
	
	prefixFinder.FindPrefix("xyz")
}