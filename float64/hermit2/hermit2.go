// The package hermit2 contains functions for 2D cubic hermit splines.
// See: http://en.wikipedia.org/wiki/Cubic_Hermite_spline
package hermit2

import (
	"fmt"
	"github.com/ungerik/go3d/float64/vec2"
)

type PointTangent struct {
	Point   vec2.T
	Tangent vec2.T
}

// T holds the data to define a hermit spline.
type T struct {
	A PointTangent
	B PointTangent
}

// Parse parses T from a string. See also String()
func Parse(s string) (r T, err error) {
	_, err = fmt.Sscanf(s,
		"%f %f %f %f %f %f %f %f",
		&r.A.Point[0], &r.A.Point[1],
		&r.A.Tangent[0], &r.A.Tangent[1],
		&r.B.Point[0], &r.B.Point[1],
		&r.B.Tangent[0], &r.B.Tangent[1],
	)
	return r, err
}

// String formats T as string. See also Parse().
func (self *T) String() string {
	return fmt.Sprintf("%s %s %s %s",
		self.A.Point.String(), self.A.Tangent.String(),
		self.B.Point.String(), self.B.Tangent.String(),
	)
}

// Point returns a point on a hermit spline at t (0,1).
func (self *T) Point(t float64) vec2.T {
	return Point(&self.A.Point, &self.A.Tangent, &self.B.Point, &self.B.Tangent, t)
}

// Tangent returns a tangent on a hermit spline at t (0,1).
func (self *T) Tangent(t float64) vec2.T {
	return Tangent(&self.A.Point, &self.A.Tangent, &self.B.Point, &self.B.Tangent, t)
}

// Length returns the length of a hermit spline from A.Point to t (0,1).
func (self *T) Length(t float64) float64 {
	return Length(&self.A.Point, &self.A.Tangent, &self.B.Point, &self.B.Tangent, t)
}

// Point returns a point on a hermit spline at t (0,1).
func Point(pointA, tangentA, pointB, tangentB *vec2.T, t float64) vec2.T {
	t2 := t * t
	t3 := t2 * t

	f := 2*t3 - 3*t2 + 1
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + t
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pAf := pointB.Scaled(f)
	result.Add(&pAf)

	return result
}

// Tangent returns a tangent on a hermit spline at t (0,1).
func Tangent(pointA, tangentA, pointB, tangentB *vec2.T, t float64) vec2.T {
	t2 := t * t
	t3 := t2 * t

	f := 2*t3 - 3*t2
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + 1
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pAf := pointB.Scaled(f)
	result.Add(&pAf)

	return result
}

// Length returns the length of a hermit spline from pointA to t (0,1).
func Length(pointA, tangentA, pointB, tangentB *vec2.T, t float64) float64 {
	sqrT := t * t
	t1 := sqrT * 0.5
	t2 := sqrT * t * 1.0 / 3.0
	t3 := sqrT*sqrT + 1.0/4.0

	f := 2*t3 - 3*t2 + t
	result := pointA.Scaled(f)

	f = t3 - 2*t2 + t1
	tAf := tangentA.Scaled(f)
	result.Add(&tAf)

	f = t3 - t2
	tBf := tangentB.Scaled(f)
	result.Add(&tBf)

	f = -2*t3 + 3*t2
	pBf := pointB.Scaled(f)
	result.Add(&pBf)

	return result.Length()
}