package heap

import "testing"

func TestMedianFinder_GetMedian(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		median  float64
		wantErr bool
	}{
		{
			name:    "test-1",
			nums:    []int{-1, 1},
			median:  0,
			wantErr: false,
		},
		{
			name:    "test-2",
			nums:    []int{1, 1},
			median:  1,
			wantErr: false,
		},
		{
			name:    "test-3",
			nums:    []int{5, 6, 7, 8, 9},
			median:  7,
			wantErr: false,
		},
		{
			name:    "test-4",
			nums:    []int{5, 6, 8, 8, 9, 10},
			median:  8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := Constructor()
			for _, num := range tt.nums {
				mf.AddNum(num)
			}
			got := mf.GetMedian()
			if got != tt.median {
				t.Errorf("MedianFinder.GetMedian() = %v, want %v", got, tt.median)
			}
		})
	}
}
