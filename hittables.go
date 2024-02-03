package raytracer

type Hittables struct {
	Objects []Hittable
}

func NewHittables(objs ...Hittable) *Hittables {
	h := Hittables{}

	if len(objs) > 0 {
		h.Objects = objs
	} else {
		h.Clear()
	}

	return &h
}

func (h *Hittables) Add(obj Hittable) {
	h.Objects = append(h.Objects, obj)
}

func (h *Hittables) Clear() {
	h.Objects = make([]Hittable, 0)
}

func (hs *Hittables) Hit(r Ray, tmin, tmax float64) (Hit, bool) {
	var hit Hit
	hitAny := false
	closest := tmax

	for i := 0; i < len(hs.Objects); i++ {
		if h, ok := hs.Objects[i].Hit(r, tmin, closest); ok {
			hit = h
			hitAny = true
			closest = h.T
		}
	}

	return hit, hitAny
}
