package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// Step 0. Jump to TestFirstSuite and then continue below
// Step 3. Start individual test suites
type MyFirstSuite struct {
	suite.Suite
}

type MySecondSuite struct {
	suite.Suite
}

func TestFirstSuite(t *testing.T) {
	// Step 1. Type "Run" and select suite.Run from the list
	// Step 2. Continue typing new(First) to get to new(MyFirstSuite)
	suite.Run(t, new(MyFirstSuite))
}

func TestSecondSuite(t *testing.T) {
	suite.Run(t, new(MySecondSuite))
}

func (s *MyFirstSuite) TestSubTests() {
	// Step 4. Rerun a subtest without the need to rerun the entire top-level test
	s.Run("first", func() {
		s.T().Log("first subTest in MyFirstSuite")
	})

	s.Run("second", func() {
		s.T().Log("first subTest in MyFirstSuite")
	})
}

// Step 5. Run test with the same method from different suites
func (s *MySecondSuite) TestSubTests() {
	s.Run("first", func() {
		s.T().Log("first subTest in MySecondSuite")
	})

	s.Run("second", func() {
		s.T().Log("first subTest in MySecondSuite")
	})
}

func (s *MySecondSuite) TestSecondTest() {
	// Step 6. Type assert.Equal(s.T(), "Hello!", "Hello!")

}
