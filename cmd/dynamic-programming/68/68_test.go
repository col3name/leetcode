package _69

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {

	outputss := []string{"Give  me  my  Romeo; and,", "when  he  shall die, Take", "him  and  cut  him out in", "little stars, And he will", "make  the  face of heaven", "so   fine   That  all the", "world  will  be  in  love", "with  night  And  pay  no", "worship   to   the garish", "sun.                     "}
	expected := []string{"Give  me  my  Romeo; and,", "when  he  shall die, Take", "him  and  cut  him out in", "little stars, And he will", "make  the  face of heaven", "so   fine  That  all  the", "world  will  be  in  love", "with  night  And  pay  no", "worship   to  the  garish", "sun.                     "}
	for i, row := range expected {
		assert.Equal(t, row, outputss[i])
	}
}

type input struct {
	maxWidth int
	words    []string
}

func Test2(t *testing.T) {
	maxWidth := 25
	words := []string{"Give", "me", "my", "Romeo;", "and,", "when", "he", "shall", "die,", "Take", "him", "and", "cut", "him", "out", "in", "little", "stars,", "And", "he", "will", "make", "the", "face", "of", "heaven", "so", "fine", "That", "all", "the", "world", "will", "be", "in", "love", "with", "night", "And", "pay", "no", "worship", "to", "the", "garish", "sun."}
	justify := fullJustify(words, maxWidth)
	fmt.Println(justify)
}

func Test(t *testing.T) {
	tests := []struct {
		name     string
		input    input
		expected []string
	}{
		{
			name: "2",
			input: input{
				maxWidth: 6,
				words:    []string{"Listen", "to", "many,", "speak", "to", "a", "few."},
			},
			expected: []string{
				"Listen", "to    ", "many, ", "speak ", "to   a", "few.  ",
			},
		},
		{
			name: "2",
			input: input{
				maxWidth: 10,
				words:    []string{"tushay", "roy", "likes", "to", "code"},
			},
			expected: []string{
				"tushay roy",
				"likes   to",
				"code      ",
			},
		},
		{
			name: "2",
			input: input{
				maxWidth: 16,
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			},
			expected: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name: "3",
			input: input{
				maxWidth: 16,
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			},
			expected: []string{"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
		{
			name: "4",
			input: input{
				maxWidth: 20,
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			},
			expected: []string{"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
		{
			name: "1",
			input: input{
				maxWidth: 25,
				words:    []string{"Give", "me", "my", "Romeo;", "and,", "when", "he", "shall", "die,", "Take", "him", "and", "cut", "him", "out", "in", "little", "stars,", "And", "he", "will", "make", "the", "face", "of", "heaven", "so", "fine", "That", "all", "the", "world", "will", "be", "in", "love", "with", "night", "And", "pay", "no", "worship", "to", "the", "garish", "sun."},
			},
			expected: []string{"Give  me  my  Romeo; and,",
				"when  he  shall die, Take",
				"him  and  cut  him out in",
				"little stars, And he will",
				"make  the  face of heaven",
				"so   fine  That  all  the",
				"world  will  be  in  love",
				"with  night  And  pay  no",
				"worship   to  the  garish",
				"sun.                     ",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := test.input
			outputss := fullJustify(in.words, in.maxWidth)
			for i, row := range test.expected {
				assert.Equal(t, row, outputss[i])
			}
		})
	}
}
