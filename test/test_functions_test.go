package assignment

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
			math.MaxUint32, 1 => 0, true
	*/

	test_uint32_true := []struct {
		x      uint32
		y      uint32
		result uint32
		flag   bool
	}{
		{42, math.MaxUint32, 41, true},
		{4294967290, 6, 0, true},
		{4294967290, 10, 4, true},
		{math.MaxUint32, 1, 0, true},
	}

	test_uint32_false := []struct {
		x      uint32
		y      uint32
		result uint32
		flag   bool
	}{
		{4294967290, 5, 4294967295, false},
		{1, 1, 2, false},
		{42, 2701, 2743, false},
	}

	for _, ts := range test_uint32_true {
		t.Run(fmt.Sprint(ts.result), func(t *testing.T) {

			sum, overflow := AddUint32(ts.x, ts.y)
			assert.Equal(t, ts.result, sum)
			assert.True(t, ts.flag, overflow)
		})
	}

	for _, ts := range test_uint32_false {
		t.Run(fmt.Sprint(ts.result), func(t *testing.T) {

			sum, overflow := AddUint32(ts.x, ts.y)
			assert.Equal(t, ts.result, sum)
			assert.False(t, ts.flag, overflow)
		})
	}
}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/

	test_numbers := []struct {
		number      float64
		ceil_number float64
	}{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}

	for _, ts := range test_numbers {
		t.Run(fmt.Sprint(ts.ceil_number), func(t *testing.T) {

			result := CeilNumber(ts.number)

			assert.Equal(t, fmt.Sprint(ts.ceil_number), fmt.Sprint(result))
		})
	}
	point := CeilNumber(42.42)

	assert.Equal(t, 42.50, point)
}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/

	test_strings := []struct {
		text   string
		sorted string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for _, ts := range test_strings {
		t.Run(ts.sorted, func(t *testing.T) {

			result := AlphabetSoup(ts.text)

			assert.Equal(t, ts.sorted, result)
		})
	}

}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/
	test_strings := []struct {
		text        string
		index       uint
		cryptedText string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 5, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}

	for _, ts := range test_strings {
		t.Run(ts.cryptedText, func(t *testing.T) {

			result := StringMask(ts.text, ts.index)

			assert.Equal(t, ts.cryptedText, result)
		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/

	result := WordSplit([2]string{"", words})

	assert.Equal(t, "not possible", result)
}

func TestVariadicSet(t *testing.T) {
	/*
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )
		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/

	set := VariadicSet([]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"})

	assert.Equal(t, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}, set)
}
