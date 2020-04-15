# Algorithms

Generic algorithms written in Go.

- Install [genny](https://github.com/cheekybits/genny) `go get github.com/cheekybits/genny`

**Merge Sort** 

```bash
wget -q -O - "https://github.com/pierDipi/go/tree/master/common/algorithms/sort_merge.go" | genny gen "T=<YOUR_TYPE>" >> your_file_name.go
```