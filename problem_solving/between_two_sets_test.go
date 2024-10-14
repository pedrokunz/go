package problem_solving

import "testing"

func TestFindCompatibleVersions(t *testing.T) {
	testCases := []struct {
		name                  string
		minDependencyVersions []int
		maxTestedVersions     []int
		compatibleVersions    CompatibleVersions
	}{
		{
			name:                  "Test case 1",
			minDependencyVersions: []int{2, 4},
			maxTestedVersions:     []int{16, 32, 96},
			compatibleVersions: CompatibleVersions{
				Count:    3,
				Versions: []int{4, 8, 16},
			},
		},
		{
			name:                  "Test case 2",
			minDependencyVersions: []int{3, 6},
			maxTestedVersions:     []int{36, 72},
			compatibleVersions: CompatibleVersions{
				Count:    4,
				Versions: []int{6, 12, 18, 36},
			},
		},
		{
			name:                  "Test case 3",
			minDependencyVersions: []int{2, 3},
			maxTestedVersions:     []int{12, 18},
			compatibleVersions: CompatibleVersions{
				Count:    1,
				Versions: []int{6},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			compatibleVersions := FindCompatibleVersions(tc.minDependencyVersions, tc.maxTestedVersions)
			if compatibleVersions.Count != tc.compatibleVersions.Count {
				t.Errorf("Expected %v, got %v", tc.compatibleVersions, compatibleVersions)
			}

			for i := range compatibleVersions.Versions {
				if compatibleVersions.Versions[i] != tc.compatibleVersions.Versions[i] {
					t.Errorf("Expected %v, got %v", tc.compatibleVersions, compatibleVersions)
				}
			}
		})
	}
}
