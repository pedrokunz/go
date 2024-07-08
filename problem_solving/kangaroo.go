package problem_solving

func Kangaroo(
	kangaroo1StartPosition,
	kangaroo1Speed,
	kangaroo2StartPosition,
	kangaroo2Speed int32,
) string {
	if kangaroo1Speed == kangaroo2Speed {
		if kangaroo1StartPosition == kangaroo2StartPosition {
			return "YES"
		}

		return "NO"
	}

	positionDifference := kangaroo2StartPosition - kangaroo1StartPosition
	speedDifference := kangaroo1Speed - kangaroo2Speed
	if positionDifference%speedDifference == 0 &&
		positionDifference/speedDifference > 0 {
		return "YES"
	}

	return "NO"
}
