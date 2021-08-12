package main


import (
	"image/color"
)


type geometry interface{
	collide_ray(R ray) (dist float64,p0 float64,p1 float64)
	get_color(tex texture, p0 float64, p1 float64 ) color.RGBA
}


type btex interface{
	color(vel vec3) color.RGBA
}
