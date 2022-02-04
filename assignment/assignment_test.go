package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/
	cases := []struct {
		valueFirst     uint32
		valueSecond    uint32
		sumResult      uint32
		overFlowResult bool
		caseNames      string
	}{
		{math.MaxUint32, 1, 0, true, "1"},
		{1, 1, 2, false, "2"},
		{42, 2701, 2743, false, "3"},
		{42, math.MaxUint32, 41, true, "4"},
		{4294967290, 5, 4294967295, false, "5"},
		{4294967290, 6, 0, true, "6"},
		{4294967290, 10, 4, true, "7"},
	}
	for _, c := range cases {
		t.Run(c.caseNames, func(t *testing.T) {
			sum, overflow := AddUint32(c.valueFirst, c.valueSecond)
			assert.Equal(t, c.sumResult, sum)
			assert.Equal(t, c.overFlowResult, overflow)
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
	cases := []struct {
		value     float64
		result    float64
		caseNames string
	}{
		{42.42, 42.50, "1"},
		{42, 42, "2"},
		{42.01, 42.25, "3"},
		{42.24, 42.25, "4"},
		{42.25, 42.25, "5"},
		{42.26, 42.50, "6"},
		{42.55, 42.75, "7"},
		{42.75, 42.75, "8"},
		{42.76, 43, "9"},
		{42.99, 43, "10"},
		{43.13, 43.25, "11"},
		{-43.46, -43.25, "negative ceil -1"},
		{-42.42, -42.25, "-2"},
		{-42.26, -42.25, "-3"},
		{-42.25, -42.25, "-4"},
		{-42.24, -42, "-5"},
		{-42.01, -42, "-6"},
		{-42, -42, "-7"},
	}
	for _, c := range cases {
		t.Run(c.caseNames, func(t *testing.T) {
			result := CeilNumber(c.value)
			assert.Equal(t, c.result, result)
		})
	}

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
	cases := []struct {
		value    string
		result   string
		caseName string
	}{
		{"hello", "ehllo", "order all char"},
		{"", "", "order empty string "},
		{"h", "h", "order one string"},
		{"ab", "ab", "order ordered string"},
		{"ba", "ab", "order standard string v1"},
		{"bac", "abc", "order standard string v2"},
		{"cba", "abc", "order standard string v3"},
		{"faik rocks", " acfikkors", "space first"},
		{"test", "estt", "same chars with same order"},
		{"faik yıldırım", " adfiklmryııı", "same chars with same order with space"},
	}
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			result := AlphabetSoup(c.value)
			assert.Equal(t, c.result, result)
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
	cases := []struct {
		value    string
		count    uint
		result   string
		caseName string
	}{
		{"!mysecret*", 2, "!m********", "standard mask(no spec)"},
		{"", 2, "*", "empty string returns one char masked"},
		{"a", 1, "*", "one char returns  one char masked "},
		{"a", 6, "*", "one char returns  one char masked v1 "},
		{"a", 0, "*", "one char with 0 returns  one char masked "},
		{"string", 0, "******", "when count 0 returns all chars masked"},
		{"string", 3, "str***", "standard mask(no special Char)"},
		{"string", 5, "strin*", "standard mask(no special Char) to given len"},
		{"string", 6, "******", "standard mask to length size"},
		{"string", 7, "******", "if bigger then len all chars masked"},
		{"faikyıldırım", 166, "************", "if bigger then len all chars masked v1"},
		{"faik yıldırım", 4,  "faik*********", "with space"},
		{"s*r*n*", 3, "s*r***", "mask no matter text include(mask char)"},
	}
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			result := StringMask(c.value, c.count)
			assert.Equal(t, c.result, result)
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
	cases := []struct {
		values   [2]string
		result   string
		caseName string
	}{
		{[2]string{"hellocat", words}, "hello,cat", "function really work"},

		{[2]string{"catbat", words}, "cat,bat", "function work for all element not just first "},
		{[2]string{"yellowapple", words}, "yellow,apple", "yes it's really need work"},
		{[2]string{"", words}, "not possible", "it's not possible on empty string"},
		{[2]string{"notcat", words}, "not possible", "also it's not on possible just one element"},
		{[2]string{"bootcamprocks!", words}, "not possible", "or no element"},
		{[2]string{"appleyellow", words}, "yellow,apple", "same text returns same order"},
		{[2]string{"cathello", words}, "hello,cat", "same text returns same order v1"},
		{[2]string{"whyapple", words}, "why,apple", "same text returns same order v2"},
		{[2]string{"applewhy", words}, "why,apple", "same text returns same order v3"},
		{[2]string{"catgoodbye", words}, "goodbye,cat", "same text returns same order v4"},
	}
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			result := WordSplit(c.values)
			assert.Equal(t, c.result, result)
		})
	}

}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/
	cases := []struct {
		values   []interface{}
		result   []interface{}
		caseName string
	}{
		{
			values:   []interface{}{4, 2, 5, 4, 2, 4},
			result:   []interface{}{4, 2, 5},
			caseName: "integer case",
		},
		{
			values:   []interface{}{"bootcamp", "rocks!", "really", "rocks!"},
			result:   []interface{}{"bootcamp", "rocks!", "really"},
			caseName: "some string",
		},
		{
			values:   []interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"},
			result:   []interface{}{1, uint32(1), "first", 2, uint32(2), "second"},
			caseName: "weird case",
		},
		{
			values:   []interface{}{"bootcamp", "rocks!", "really", "aaa!", "bbbb", "cccc", "dddd", "www", "qqqq"},
			result:   []interface{}{"bootcamp", "rocks!", "really", "aaa!", "bbbb", "cccc", "dddd", "www", "qqqq"},
			caseName: "all different string ",
		},
		{
			values:   []interface{}{"bootcamp", "bootcamp", "bootcamp", "bootcamp", "bootcamp", "bootcamp", "bootcamp", "bootcamp", "bootcamp"},
			result:   []interface{}{"bootcamp"},
			caseName: " all same string",
		},
		{
			values: []interface{}{
				struct {
					test string
					intValue int
				}{"try this",2},
				struct {
					test string
					intValue int
				}{"try this",2},
				struct {
					test string
				}{"not this"},
				struct {
					testInt int
				}{12},
			},

			result: []interface{}{
				struct {
					test string
					intValue int
				}{"try this",2},
				struct {
					test string
				}{"not this"},
				struct {
					testInt int
				}{12},
			},
			caseName: " what if it struct",
		},

	}
	for _, c := range cases {
		t.Run(c.caseName, func(t *testing.T) {
			set := VariadicSet(c.values...)
			assert.Equal(t, c.result, set)
		})
	}

}
