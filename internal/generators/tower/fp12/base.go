package fp12

const Base = `
import (
	"github.com/consensys/gurvy/{{.Fpackage}}/fp"
)

// {{.Fp12Name}} is a degree-two finite field extension of fp6:
// C0 + C1w where w^3-v is irrep in fp6

// fp2, fp12 are both quadratic field extensions
// template code is duplicated in fp2, fp12
// TODO make an abstract quadratic extension template

type {{.Fp12Name}} struct {
	C0, C1 {{.Fp6Name}}
}

// Equal returns true if z equals x, fasle otherwise
// TODO can this be deleted?  Should be able to use == operator instead
func (z *{{.Fp12Name}}) Equal(x *{{.Fp12Name}}) bool {
	return z.C0.Equal(&x.C0) && z.C1.Equal(&x.C1)
}

// String puts {{.Fp12Name}} in string form
func (z *{{.Fp12Name}}) String() string {
	return (z.C0.String() + "+(" + z.C1.String() + ")*w")
}

// SetString sets a {{.Fp12Name}} from string
func (z *{{.Fp12Name}}) SetString(s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11 string) *{{.Fp12Name}} {
	z.C0.SetString(s0, s1, s2, s3, s4, s5)
	z.C1.SetString(s6, s7, s8, s9, s10, s11)
	return z
}

// Set copies x into z and returns z
func (z *{{.Fp12Name}}) Set(x *{{.Fp12Name}}) *{{.Fp12Name}} {
	z.C0 = x.C0
	z.C1 = x.C1
	return z
}

// SetOne sets z to 1 in Montgomery form and returns z
func (z *{{.Fp12Name}}) SetOne() *{{.Fp12Name}} {
	z.C0.B0.A0.SetOne()
	z.C0.B0.A1.SetZero()
	z.C0.B1.A0.SetZero()
	z.C0.B1.A1.SetZero()
	z.C0.B2.A0.SetZero()
	z.C0.B2.A1.SetZero()
	z.C1.B0.A0.SetZero()
	z.C1.B0.A1.SetZero()
	z.C1.B1.A0.SetZero()
	z.C1.B1.A1.SetZero()
	z.C1.B2.A0.SetZero()
	z.C1.B2.A1.SetZero()
	return z
}

// ToMont converts to Mont form
// TODO can this be deleted?
func (z *{{.Fp12Name}}) ToMont() *{{.Fp12Name}} {
	z.C0.ToMont()
	z.C1.ToMont()
	return z
}

// FromMont converts from Mont form
// TODO can this be deleted?
func (z *{{.Fp12Name}}) FromMont() *{{.Fp12Name}} {
	z.C0.FromMont()
	z.C1.FromMont()
	return z
}

// Add set z=x+y in {{.Fp12Name}} and return z
func (z *{{.Fp12Name}}) Add(x, y *{{.Fp12Name}}) *{{.Fp12Name}} {
	z.C0.Add(&x.C0, &y.C0)
	z.C1.Add(&x.C1, &y.C1)
	return z
}

// Sub set z=x-y in {{.Fp12Name}} and return z
func (z *{{.Fp12Name}}) Sub(x, y *{{.Fp12Name}}) *{{.Fp12Name}} {
	z.C0.Sub(&x.C0, &y.C0)
	z.C1.Sub(&x.C1, &y.C1)
	return z
}

// SetRandom used only in tests
// TODO eliminate this method!
func (z *{{.Fp12Name}}) SetRandom() *{{.Fp12Name}} {
	z.C0.B0.A0.SetRandom()
	z.C0.B0.A1.SetRandom()
	z.C0.B1.A0.SetRandom()
	z.C0.B1.A1.SetRandom()
	z.C0.B2.A0.SetRandom()
	z.C0.B2.A1.SetRandom()
	z.C1.B0.A0.SetRandom()
	z.C1.B0.A1.SetRandom()
	z.C1.B1.A0.SetRandom()
	z.C1.B1.A1.SetRandom()
	z.C1.B2.A0.SetRandom()
	z.C1.B2.A1.SetRandom()
	return z
}
`
