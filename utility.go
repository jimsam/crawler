package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func checkAndCastArguments(args []string) (bool, []int, error) {
	switch len(args) {
	case 0:
		return false, []int{}, errors.New("no website provided")
	case 1, 2:
		return false, []int{}, errors.New("too few arguments")
	case 3:
		_, argSlice, err := validArguments(args)
		if err != nil {
			return false, []int{}, err
		}
		return true, argSlice, nil
	default:
		return false, []int{}, errors.New("too many arguments")
	}
}

func validArguments(args []string) (bool, []int, error) {
	_, arg1, err := CheckStringIsInt(args[1])
	if err != nil {
		return false, []int{}, fmt.Errorf("There was an error with concurrency arg: %w", err)
	}

	_, arg2, err := CheckStringIsInt(args[2])
	if err != nil {
		return false, []int{}, fmt.Errorf("There was an error with max pages arg: %w", err)
	}
	return true, []int{arg1, arg2}, nil
}

func CheckStringIsInt(str string) (bool, int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return false, 0, err
	}
	return true, num, nil
}

func SortMapByIntVal(hashMap map[string]int, asc bool) []Page {
	direction := 1
	if asc == false {
		direction = -1
	}

	res := []Page{}
	for key, val := range hashMap {
		res = append(res, Page{url: key, count: val})
	}

	slices.SortFunc(res, func(i, j Page) int {
		if n := cmp.Compare(i.count, j.count) * direction; n != 0 {
			return n
		}

		return strings.Compare(i.url, j.url) * direction
	})
	return res
}
