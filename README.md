This package contains implementation of 3 dynamic programming algorithms:

- [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance) (edit distance)
- [Needleman-Wunsch algorithm](https://en.wikipedia.org/wiki/Needleman%E2%80%93Wunsch_algorithm) (sequence alignment)
- [Hirschberg algorithm](https://en.wikipedia.org/wiki/Hirschberg%27s_algorithm) (space-optimised NW algorithm)
- [Longest common subsequence](https://en.wikipedia.org/wiki/Longest_common_subsequence_problem)

### Levenshtein distance

```go
d := LevenshteinDistance([]byte("kitten"), []byte("sitting"), nil, nil)
//d ==3
```

## Needleman-Wunsch
The NW implementation requires the user to provide a similarity function that can return the similarity between two points
```go
func similarity(a byte, b byte) int {
	if a == b {
		return 1
	}
	return -1
}
a := []byte("ABCDEF")
b := []byte("ABCCDEF")

z, w := NeedlemanWunsch(a, b, -1, similarity, f)

```

