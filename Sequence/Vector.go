package Sequence

import "errors"

type Vector interface {
	getSize() int
	isEmpty() bool
	getAtRank(rank int) (string, error)
	replaceAtRank(rank int, str string)(bool, error)
	insertAtRank(rank int, str string)(bool, error)
	removeAtRank(rank int)(bool, error)
}

type VectorArray struct {
	cap int
	size int
	data []string
}

func NewVectoryArray() *VectorArray {
	return &VectorArray{
		cap:  1024,
		size: 0,
		data: make([]string, 0),
	}
}

func (v *VectorArray) getSize() int {
	return v.size
}

func (v *VectorArray) isEmpty() bool {
	return v.size == 0
}

func (v *VectorArray) getAtRank(rank int) (string, error)  {
	if 0 > rank || rank >= v.size {
		return "", errors.New("Exception OverflowRank")
	}
	return v.data[rank], nil
}

func (v *VectorArray) replaceAtRank(rank int, str string)(bool, error)  {
	if 0 > rank || rank >= v.size {
		return false, errors.New("Exception OverflowRank")
	}
	return true, nil
}