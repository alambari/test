package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// GO1 mengembalikan angka yang "hilang" dari arr:
//   - campuran negatif & positif -> 0
//   - hanya positif              -> bilangan positif terkecil yang tidak ada (1,2,3,...)
//   - hanya negatif              -> bilangan negatif terbesar yang tidak ada (-1,-2,-3,...)
//
// Angka 0 diabaikan. Kompleksitas O(n) memakai set.
func GO1(arr []int) int {
	seen := make(map[int]struct{}, len(arr))
	hasNeg, hasPos := false, false

	// Satu kali lewat: catat semua angka dan tandai ada/tidaknya positif & negatif.
	for _, v := range arr {
		switch {
		case v > 0:
			hasPos = true
		case v < 0:
			hasNeg = true
		default:
			continue // 0 diabaikan
		}
		seen[v] = struct{}{}
	}

	// Campuran (atau kosong/semua nol) -> 0.
	if hasNeg == hasPos {
		return 0
	}

	// Hanya positif: cari 1,2,3,... pertama yang tidak ada. Hanya negatif: arah dibalik.
	step := 1
	if hasNeg {
		step = -1
	}
	for n := step; ; n += step {
		if _, ok := seen[n]; !ok {
			return n
		}
	}
}

// GO2 menggabungkan nums menjadi satu string digit, membaliknya, lalu membuang
// digit yang ada di exclude. Hasil kosong -> (0, "not found").
func GO2(exclude []int, nums ...int) (int, error) {
	// Set digit terlarang agar pengecekan O(1).
	skip := make(map[byte]struct{}, len(exclude))
	for _, e := range exclude {
		skip[byte('0'+e)] = struct{}{}
	}

	// Gabungkan semua angka menjadi satu deret digit.
	var joined strings.Builder
	for _, n := range nums {
		joined.WriteString(strconv.Itoa(n))
	}
	s := joined.String()

	// Baca dari belakang (= reverse) sambil melewati digit yang di-exclude.
	var out strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if _, bad := skip[s[i]]; !bad {
			out.WriteByte(s[i])
		}
	}
	if out.Len() == 0 {
		return 0, errors.New("not found")
	}

	// Atoi otomatis membuang leading zero ("01" -> 1).
	res, err := strconv.Atoi(out.String())
	if err != nil {
		return 0, err
	}
	return res, nil
}

// GO3 mengembalikan angka unik di numsA yang juga ada di numsB, urut menurun.
func GO3(numsA, numsB []int) []int {
	inB := make(map[int]struct{}, len(numsB))
	for _, v := range numsB {
		inB[v] = struct{}{}
	}

	// Ambil irisan; map "added" mencegah duplikat (mis. [1,2,2] -> 2 sekali).
	res := []int{}
	added := make(map[int]struct{})
	for _, v := range numsA {
		if _, ok := inB[v]; !ok {
			continue
		}
		if _, dup := added[v]; dup {
			continue
		}
		added[v] = struct{}{}
		res = append(res, v)
	}

	// Urut descending.
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return res
}

// GO4 memecah num menjadi digit lalu menghitung kemunculan tiap digit.
func GO4(num int) map[int]int {
	if num < 0 {
		num = -num // tanda minus bukan digit
	}
	count := make(map[int]int)

	// num == 0 tidak masuk loop, jadi ditangani khusus.
	if num == 0 {
		count[0] = 1
		return count
	}

	// Ambil digit paling kanan dengan modulo, lalu geser dengan pembagian.
	for ; num > 0; num /= 10 {
		count[num%10]++
	}
	return count
}

// GO5 mengalikan pasangan (index genap, index ganjil) secara concurrent, satu
// goroutine per pasangan, lalu menjumlahkan hasilnya di main thread.
// Elemen terakhir tanpa pasangan dikali dirinya sendiri.
func GO5(nums ...int) int {
	pairs := (len(nums) + 1) / 2
	// Buffered sebesar jumlah pasangan: goroutine tidak pernah blocking saat kirim.
	results := make(chan int, pairs)

	for i := 0; i < len(nums); i += 2 {
		a, b := nums[i], nums[i]
		if i+1 < len(nums) {
			b = nums[i+1]
		}
		go func() { results <- a * b }()
	}

	// Penjumlahan di main thread: tepat sebanyak pasangan.
	sum := 0
	for i := 0; i < pairs; i++ {
		sum += <-results
	}
	return sum
}

// section mencetak judul tiap soal.
func section(title string) {
	fmt.Printf("\n━━ %s %s\n", title, strings.Repeat("━", 44-len([]rune(title))))
}

// row mencetak satu pemanggilan fungsi beserta hasilnya dalam kolom rapi.
func row(call string, result any) {
	fmt.Printf("  %-36s → %v\n", call, result)
}

func main() {
	section("Soal 1 · angka yang hilang")
	row("GO1([-1 1])", GO1([]int{-1, 1}))
	row("GO1([-1 -7 -5])", GO1([]int{-1, -7, -5}))
	row("GO1([1 2 1 6])", GO1([]int{1, 2, 1, 6}))

	section("Soal 2 · gabung, reverse, exclude")
	n, err := GO2([]int{0, 1, 0}, 1000, 1010)
	row("GO2([0 1 0], 1000, 1010)", fmt.Sprintf("%d, error: %v", n, err))
	n, err = GO2([]int{0, 1, 0}, 1259, 2601)
	row("GO2([0 1 0], 1259, 2601)", fmt.Sprintf("%d, error: %v", n, err))

	section("Soal 3 · irisan descending")
	row("GO3([1 2], [1 3])", GO3([]int{1, 2}, []int{1, 3}))
	row("GO3([1 2 2], [1 2 4])", GO3([]int{1, 2, 2}, []int{1, 2, 4}))

	section("Soal 4 · hitung digit")
	row("GO4(1223334)", GO4(1223334))

	section("Soal 5 · perkalian concurrent")
	row("GO5(1, 2, 3, 4)", GO5(1, 2, 3, 4))
	row("GO5(1, 2, 3, 4, 5)", GO5(1, 2, 3, 4, 5))
	fmt.Println()
}
