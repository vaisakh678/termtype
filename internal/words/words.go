package words

import (
	"math/rand"
)

var common = []string{
	"the", "be", "to", "of", "and", "a", "in", "that", "have", "i",
	"it", "for", "not", "on", "with", "he", "as", "you", "do", "at",
	"this", "but", "his", "by", "from", "they", "we", "say", "her", "she",
	"or", "an", "will", "my", "one", "all", "would", "there", "their", "what",
	"so", "up", "out", "if", "about", "who", "get", "which", "go", "me",
	"when", "make", "can", "like", "time", "no", "just", "him", "know", "take",
	"people", "into", "year", "your", "good", "some", "could", "them", "see", "other",
	"than", "then", "now", "look", "only", "come", "its", "over", "think", "also",
	"back", "after", "use", "two", "how", "our", "work", "first", "well", "way",
	"even", "new", "want", "because", "any", "these", "give", "day", "most", "us",
	"great", "between", "need", "large", "often", "hand", "high", "place", "hold",
	"free", "real", "life", "few", "north", "open", "seem", "together", "next",
	"white", "children", "begin", "got", "walk", "example", "ease", "paper", "group",
	"always", "music", "those", "both", "mark", "book", "letter", "until", "mile",
	"river", "car", "feet", "care", "second", "enough", "plain", "girl", "usual",
	"young", "ready", "above", "ever", "red", "list", "though", "feel", "talk",
	"bird", "soon", "body", "dog", "family", "direct", "pose", "leave", "song",
	"measure", "door", "product", "black", "short", "number", "class", "wind",
	"question", "happen", "complete", "ship", "area", "half", "rock", "order",
	"fire", "south", "problem", "piece", "told", "knew", "pass", "since", "top",
	"whole", "king", "space", "heard", "best", "hour", "better", "true", "during",
	"hundred", "five", "remember", "step", "early", "left", "didn", "world",
	"going", "keep", "does", "machine", "start", "might", "story", "under",
	"point", "city", "run", "didn", "every", "right", "move", "still", "should",
}

// Generate returns n random words from the common word list.
func Generate(n int) []string {
	words := make([]string, n)
	for i := range words {
		words[i] = common[rand.Intn(len(common))]
	}
	return words
}
