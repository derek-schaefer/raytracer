package raytracer

type Hittables struct {
	Objects []Hittable
}

func NewHittables(objs ...Hittable) *Hittables {
	return &Hittables{objs}
}

func (hs *Hittables) Add(obj Hittable) {
	hs.Objects = append(hs.Objects, obj)
}

func (hs *Hittables) Clear() {
	hs.Objects = make([]Hittable, 0)
}

func (hs *Hittables) Hit(r Ray, rt Interval) (Hit, bool) {
	var hit Hit
	hitAny := false
	closest := rt.Max

	for i := 0; i < len(hs.Objects); i++ {
		if h, ok := hs.Objects[i].Hit(r, NewInterval(rt.Min, closest)); ok {
			hit = h
			hitAny = true
			closest = h.T
		}
	}

	return hit, hitAny
}
