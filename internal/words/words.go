package words

import (
	"math/rand"
)

var common = []string{
	// Basic / short words
	"the", "be", "to", "of", "and", "a", "in", "that", "have", "it",
	"for", "not", "on", "with", "he", "as", "you", "do", "at", "this",
	"but", "his", "by", "from", "they", "we", "say", "her", "she", "or",
	"an", "will", "my", "one", "all", "would", "there", "their", "what",
	"so", "up", "out", "if", "about", "who", "get", "which", "go", "me",
	"when", "make", "can", "like", "time", "no", "just", "him", "know", "take",
	"people", "into", "year", "your", "good", "some", "could", "them", "see", "other",
	"than", "then", "now", "look", "only", "come", "its", "over", "think", "also",
	"back", "after", "use", "two", "how", "our", "work", "first", "well", "way",
	"even", "new", "want", "because", "any", "these", "give", "day", "most", "us",

	// Common nouns
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
	"hundred", "five", "remember", "step", "early", "left", "world",
	"going", "keep", "does", "machine", "start", "might", "story", "under",
	"point", "city", "run", "every", "right", "move", "still", "should",

	// Extended words
	"house", "light", "picture", "try", "change", "off", "play", "spell",
	"air", "away", "animal", "answer", "build", "thought", "let", "found",
	"study", "still", "learn", "plant", "cover", "food", "sun", "four",
	"state", "eye", "never", "last", "door", "between", "turn", "cross",
	"start", "close", "head", "hard", "along", "line", "face", "name",
	"show", "form", "much", "mean", "move", "old", "same", "tell",
	"help", "low", "differ", "earth", "near", "add", "father", "mother",
	"land", "here", "must", "big", "draw", "went", "man", "read",
	"own", "page", "while", "press", "night", "side", "been", "call",
	"cut", "ask", "late", "end", "set", "home", "small", "hot",
	"sentence", "bring", "word", "money", "serve", "appear", "road", "map",
	"rain", "rule", "govern", "pull", "cold", "notice", "voice", "energy",
	"act", "reach", "market", "figure", "table", "travel", "morning",
	"simple", "several", "grow", "less", "write", "war", "against",
	"pattern", "slow", "center", "love", "person", "fact", "street",
	"inch", "lot", "nothing", "course", "stay", "wheel", "full", "force",
	"blue", "object", "decide", "surface", "deep", "moon", "island",
	"foot", "system", "busy", "test", "record", "boat", "common",
	"gold", "possible", "plan", "age", "wonder", "laugh", "check",
	"game", "shape", "miss", "heat", "snow", "tire", "fill",
	"east", "paint", "language", "among", "unit", "power", "town",
	"fine", "fly", "fall", "lead", "cry", "dark", "climb",
	"dream", "drive", "field", "rest", "heart", "heavy", "hope",
	"minute", "strong", "special", "mind", "behind", "clear", "tail",
	"produce", "stand", "flat", "contain", "front", "teach", "season",
	"tool", "ground", "spring", "case", "water", "sort", "rise",
	"window", "figure", "store", "summer", "train", "sleep", "prove",
	"catch", "mount", "board", "round", "swift", "block", "chart",
	"sound", "crowd", "quiet", "stone", "tiny", "track", "glass",
	"floor", "tower", "ocean", "bright", "sharp", "cloud", "master",
	"dinner", "fruit", "anger", "claim", "soft", "bread", "strange",
}

// Generate returns n random words from the common word list.
func Generate(n int) []string {
	words := make([]string, n)
	for i := range words {
		words[i] = common[rand.Intn(len(common))]
	}
	return words
}
