package banknotecalculator

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeat(t *testing.T) {
	r := NewRepeat([]uint{1, 2, 4}, uint64(maxSupply))

	for i := 0; i < 100; i++ {
		n := int64(minSupply) + rand.Int63n(int64(maxSupply-minSupply))

		used, cnt := r.Cal(uint64(n)), uint64(0)
		for unit, u := range used {
			cnt += unit * uint64(u)
		}
		require.Equal(t, cnt, uint64(n))

		fmt.Println(n, used)
	}
}

func TestRepeatBestUnits(t *testing.T) {
	bestUnits := make(map[string]int)

	for sup := uint64(minSupply); sup < uint64(maxSupply); sup += uint64(step) {
		best := supplyBestUnits(sup)
		bestUnits[unitsKey(best)]++
	}

	for units, n := range bestUnits {
		fmt.Println(units, n)
	}
}

func unitsToSupply(sup uint64, base []uint64) []uint64 {
	var units []uint64
	next, dec := 0, 1
	for {
		u := base[next] * uint64(dec)
		if u > sup {
			return units
		}

		units = append(units, u)

		if next == len(base)-1 {
			next, dec = 0, dec*10
		} else {
			next++
		}
	}
}

func supplyBestUnits(sup uint64) []uint64 {
	var best []uint64
	minTotal, bestMaxBills, maxUsed := uint64(math.MaxInt64), uint64(0), uint64(0)

	for i := uint64(2); i <= 9; i++ {
		for j := uint64(2); j <= 9; j++ {
			if i == j {
				continue
			}

			total, base, maxBills, curMaxUsed := 0, []uint64{1, i, j}, uint64(0), uint64(0)

			units := unitsToSupply(sup, base)
			slices.Reverse(units)

			for n := uint64(1); n <= sup; n++ {
				bills, used := calculateBills(n, units)
				maxBills = max(maxBills, bills)
				curMaxUsed = max(curMaxUsed, used)
				total += int(bills)
			}

			if uint64(total) < minTotal {
				minTotal = uint64(total)
				best = base
				bestMaxBills = maxBills
				maxUsed = curMaxUsed
			}
		}
	}

	fmt.Println(sup, best, minTotal/sup, minTotal, bestMaxBills, maxUsed)

	return best
}

func calculateBills(n uint64, units []uint64) (uint64, uint64) {
	bills, maxUsed := uint64(0), uint64(0)
	for _, unit := range units {
		if n == 0 {
			break
		}

		if n < unit {
			continue
		}

		u := n / unit
		maxUsed = max(maxUsed, u)
		bills += u
		n %= unit
	}

	return bills, maxUsed
}

func unitsKey(units []uint64) string {
	n := uint64(0)
	for _, v := range units {
		n = n*10 + v
	}

	return strconv.FormatUint(n, 10)
}
