package multijoinanalytics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Route high-value analytics results to the first branch and all others to the second branch.

func TestRouteAnalyticsResult_BuildSwitch(t *testing.T) {
	t.Skip("not yet implemented") // TODO: remove when implementation is ready
	f := &RouteAnalyticsResult{}
	// TODO: populate whenItems with transformation.When entries matching your case branches
	switchFn, err := f.BuildSwitch(nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, switchFn)
}
