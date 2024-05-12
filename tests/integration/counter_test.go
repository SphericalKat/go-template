package integration_test

import (
	"testing"

	genapi "github.com/goalistdeveloper/go-template/api/gen"
	"github.com/goalistdeveloper/go-template/tests"
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
	s.RNil(err)

	// ideally autogenerate the two lines above
	// use empty structs for structs in request
	// whatever comes after is manual work
	s.Equal(counter.Count, float64(0))
}

func (s *counterTestSuite) TestIncrementCount() {
	_, err := s.Client.IncrementCount(s.Ctx, genapi.NewOptIncrement(genapi.Increment{Amount: genapi.NewOptInt(5)}))
	s.RNil(err)

	counter, err := s.Client.GetCurrentCount(s.Ctx)
	s.RNil(err)

	s.Equal(counter.Count, float64(0))
}
