package twoSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type topic struct {
	nums   []int
	target int
}

type answer struct {
	result []int
}

type question struct {
	topic
	answer
}

func TestTwoSum(t *testing.T) {
	questions := []question{
		{
			topic{
				nums:   []int{2, 7, 11, 17},
				target: 9,
			},
			answer{result: []int{0, 1}},
		},
		{
			topic{
				nums:   []int{4, 6, 9},
				target: 15,
			},
			answer{
				result: []int{1, 2},
			},
		},
	}
	for _, q := range questions {
		assert.Equal(t, q.result, twoSum(q.nums, q.target), "输入：%v", q.topic)
	}
}
