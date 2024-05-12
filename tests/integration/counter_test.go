package integration_test

import (
	"testing"

	genapi "github.com/goalistdeveloper/go-template/api/gen"
	"github.com/goalistdeveloper/go-template/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestCounterTestSuite(t *testing.T) {
	suite.Run(t, new(counterTestSuite))
}

type counterTestSuite struct {
	tests.TestSuiteBase
}

func (s *counterTestSuite) TestGetCurrentCount() {
	counter, err := s.Client.GetCurrentCount(s.Ctx)
	require.Nil(s.T(), err)

	// ideally autogenerate the two lines above
	// use empty structs for structs in request
	// whatever comes after is manual work
	assert.Equal(s.T(), counter.Count, float64(0))
}

func (s *counterTestSuite) TestIncrementCount() {
	_, err := s.Client.IncrementCount(s.Ctx, genapi.NewOptIncrement(genapi.Increment{Amount: genapi.NewOptInt(5)}))
	require.Nil(s.T(), err)

	counter, err := s.Client.GetCurrentCount(s.Ctx)
	require.Nil(s.T(), err)

	assert.Equal(s.T(), counter.Count, float64(0))
}
