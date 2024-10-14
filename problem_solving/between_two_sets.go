package problem_solving

/*
	The "Between Two Sets" problem and its solution can be applied in various real-world scenarios,
	particularly in areas involving optimization, scheduling, and compatibility checks. Here are a few examples:

	Scheduling and Resource Allocation:
		In a complex project management scenario, resources such as machinery, meeting rooms, or skilled personnel are
		often limited and must be allocated efficiently among various tasks. The "Between Two Sets" approach can be used
		to optimize the schedule by finding time slots that align with the availability of all required
		resources (arrayA) and the scheduling needs of all tasks (arrayB). This ensures that resources are utilized to
		their maximum potential while minimizing idle time, leading to a more efficient project timeline.

	Software Compatibility:
		When managing a software project that relies on multiple external libraries or APIs, ensuring compatibility
		across different versions can be challenging. By treating the minimum required versions of these dependencies as
		arrayA and the maximum supported versions tested for compatibility as arrayB, the "Between Two Sets" solution
		can identify the range of versions for each dependency that ensures compatibility across the board.
		This approach simplifies dependency management and helps prevent integration issues.

	Manufacturing and Supply Chain:
		In manufacturing, ensuring that components sourced from different suppliers fit together in the final assembly
		requires precise coordination. By considering the tolerance levels of components as arrayA and the specifications
		of the assembly as arrayB, the "Between Two Sets" method can be used to identify the specifications that all
		components must meet to ensure a perfect fit. This can significantly reduce the risk of production delays and
		increase the efficiency of the supply chain.

	Educational Curriculum Planning:
		For educational institutions, designing a curriculum that allows students to progress through courses in a
		logical and efficient manner is crucial. By treating course prerequisites as arrayA and courses that require
		those prerequisites as arrayB, the "Between Two Sets" solution can help in organizing courses in a way that
		ensures all prerequisites are met before a student enrolls in an advanced course. This facilitates a smoother
		educational journey for students and can enhance learning outcomes.

	Cross-Platform Development:
		Developers working on cross-platform applications must ensure that their code runs smoothly on all target
		platforms, each with its own set of supported APIs and libraries. By considering the APIs supported by the least
		capable platform as arrayA and the requirements of the application as arrayB, the "Between Two Sets" approach
		can help identify a common set of APIs that can be used across all platforms. This ensures compatibility and
		functionality across different environments, simplifying the development process.

	These examples illustrate how the abstract mathematical solution of finding integers between two sets can be applied
	to solve practical problems in optimization, compatibility, and planning across various domains.
*/

type CompatibleVersions struct {
	Count    int
	Versions []int
}

// FindCompatibleVersions counts the number of versions that are compatible with all minimum dependency versions
// and all maximum tested versions.
func FindCompatibleVersions(minDependencyVersions []int, maxTestedVersions []int) CompatibleVersions {
	lowestCompatibleVersion := getLowestCompatibleVersion(minDependencyVersions)
	highestSupportedVersion := getHighestSupportedVersion(maxTestedVersions)
	compatibleVersion := getCompatibleVersion(lowestCompatibleVersion, highestSupportedVersion)

	return compatibleVersion
}

func getCompatibleVersion(lowestCompatibleVersion int, highestSupportedVersion int) CompatibleVersions {
	compatibleVersionCount := 0
	versions := []int{}
	for version := lowestCompatibleVersion; version <= highestSupportedVersion; version += lowestCompatibleVersion {
		isVersionFullyCompatible := highestSupportedVersion%version == 0
		if isVersionFullyCompatible {
			compatibleVersionCount++
			versions = append(versions, version)
		}
	}

	return CompatibleVersions{
		Count:    compatibleVersionCount,
		Versions: versions,
	}
}

// getLowestCompatibleVersion Calculate the lowest compatible version of all minimum dependency versions
func getLowestCompatibleVersion(minDependencyVersions []int) int {
	lowestCompatibleVersion := minDependencyVersions[0]
	for _, version := range minDependencyVersions[1:] {
		lowestCompatibleVersion = calculateLCM(lowestCompatibleVersion, version)
	}

	return lowestCompatibleVersion
}

// getHighestSupportedVersion Calculate the highest supported version of all maximum tested versions
func getHighestSupportedVersion(maxTestedVersions []int) int {
	highestSupportedVersion := maxTestedVersions[0]
	for _, version := range maxTestedVersions[1:] {
		highestSupportedVersion = calculateGCD(highestSupportedVersion, version)
	}

	return highestSupportedVersion
}

// calculateLCM Calculate the Least Common Multiple (LCM) of two numbers using GCD.
func calculateLCM(firstNumber, secondNumber int) int {
	return firstNumber / calculateGCD(firstNumber, secondNumber) * secondNumber
}

// calculateGCD Calculate the Greatest Common Divisor (GCD) of two numbers.
func calculateGCD(firstNumber, secondNumber int) int {
	if secondNumber == 0 {
		return firstNumber
	}

	return calculateGCD(secondNumber, firstNumber%secondNumber)
}
