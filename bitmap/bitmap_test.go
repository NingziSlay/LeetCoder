package main

import "testing"

func TestNewBitmap(t *testing.T) {
	var b = NewBitmap(0)
	if len(b.bitmap) != 0 {
		t.Fatal("length of bitmap should be 0")
	}
}

func TestBitmap_Add(t *testing.T) {
	var b = NewBitmap(0)
	cases := []struct {
		input  uint32
		index  int
		expect uint32
		length int
	}{
		{input: 1, index: 0, expect: 2, length: 1},
		{input: 2, index: 0, expect: 6, length: 1},
		{input: 31, index: 0, expect: 2147483654, length: 1},
		{input: 32, index: 1, expect: 1, length: 2},
		{input: 63, index: 1, expect: 2147483649, length: 2},
	}

	for _, c := range cases {
		b.Add(c.input)
		if len(b.bitmap) != c.length {
			t.Fatal("length of bitmap should automatic expansion")
		}
		if b.bitmap[c.index] != c.expect {
			t.Fatal("unexpected value")
		}
	}
}

func TestBitmap_Contain(t *testing.T) {
	var b = NewBitmap(0)
	var input uint32 = 1
	var notInput uint32 = 2
	b.Add(input)
	if !b.Contain(input) {
		t.Fatalf("bitmap should contain item %d", input)
	}
	if b.Contain(notInput) {
		t.Fatalf("bitmap should not contaim item %d", notInput)
	}
}

func TestBitmap_Del(t *testing.T) {
	var b = NewBitmap(0)
	var input uint32 = 1
	b.Add(input)
	b.Del(input)
	if b.Contain(input) {
		t.Fatalf("bitmap should not contain item %d", input)
	}
}

func TestBitmap_String(t *testing.T) {
	var b = NewBitmap(0)
	if b.String() != "" {
		t.Fatal("error output String()")
	}
	b.Add(3)
	if b.String() != "00000000000000000000000000001000" {
		t.Fatal("error output String()")
	}
	b.Add(1)
	if b.String() != "00000000000000000000000000001010" {
		t.Fatal("error output String()")
	}
	b.Add(32)
	if b.String() != "0000000000000000000000000000000100000000000000000000000000001010" {
		t.Fatal("error output String()")
	}
}

func BenchmarkBitmap_Add(b *testing.B) {
	bm := NewBitmap(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bm.Add(uint32(i))
	}
}
