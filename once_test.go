package banknotecalculator

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	minSupply, maxSupply = 10000, 1000000000
	step                 = minSupply
)

func TestOnce(t *testing.T) {
	_, ok := NewOnce([]uint{1, 2, 3})
	require.False(t, ok)

	o, ok := NewOnce([]uint{1, 2, 3, 6})
	require.True(t, ok)

	for i := 0; i < 100; i++ {
		n := int64(minSupply) + rand.Int63n(int64(maxSupply-minSupply))

		used, cnt := o.Cal(uint64(n)), uint64(0)
		for _, u := range used {
			cnt += u
		}
		require.Equal(t, cnt, uint64(n))

		fmt.Println(n, used)
	}
}

func TestOnceBestUnits(t *testing.T) {
	for i := 2; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				units := []int{1, i, j, k}

				var useOnce func(n int, v []bool) int
				useOnce = func(n int, v []bool) int {
					if n == 0 {
						return 0
					}
					res := math.MaxInt32
					for i := 0; i < 4; i++ {
						if v[i] || n < units[i] {
							continue
						}
						v[i] = true
						nx := useOnce(n-units[i], v)
						if nx >= 0 {
							res = min(res, 1+nx)
						}
						v[i] = false
					}
					if res == math.MaxInt32 {
						return -1
					}
					return res
				}

				canUseOnce, useds := true, make(map[int]int)
				for n := 1; n <= 9; n++ {
					if used := useOnce(n, make([]bool, 4)); used > 0 {
						useds[n] = used
					} else {
						canUseOnce = false
						break
					}
				}

				if canUseOnce {
					total := 0
					for _, used := range useds {
						total += used
					}

					fmt.Println(units, total, useds)
				}
			}
		}
	}
}
