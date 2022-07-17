package geom

import "goudptest/internal/util/binpack"

type Vec3 struct {
	X, Y, Z float32
}

func (v *Vec3) Add(dv Vec3) {
	v.X += dv.X
	v.Y += dv.Y
	v.Z += dv.Z
}
func (v1 Vec3) Plus(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func (v1 Vec3) Minus(v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func (v Vec3) SizeSquared() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vec3) Marshall(bp *binpack.Binpacker) {
	bp.MarshallFloat32(v.X)
	bp.MarshallFloat32(v.Y)
	bp.MarshallFloat32(v.Z)
}
func (v *Vec3) Unmarshall(bp *binpack.Binpacker) error {
	x, xerr := bp.UnmarshallFloat32()
	if xerr != nil {
		return xerr
	}

	y, yerr := bp.UnmarshallFloat32()
	if yerr != nil {
		return yerr
	}

	z, zerr := bp.UnmarshallFloat32()
	if zerr != nil {
		return zerr
	}

	v.X = x
	v.Y = y
	v.Z = z

	return nil
}
