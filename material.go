package raytracer

type Material interface {
	Scatter(in Ray, hit Hit) (Ray, Color, bool)
}
