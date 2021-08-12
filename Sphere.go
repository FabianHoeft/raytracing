package main


import (
	"math"
	"image/color"
)

type sphere struct {
  pos vec3
  radius float64
  color color.RGBA
}

func (S sphere) collide_ray( R ray) (float64,float64,float64) {
  temp:= vec3{R.pos.x-S.pos.x,R.pos.y-S.pos.y,R.pos.z-S.pos.z}
  delta:=math.Pow(dot(R.vel,temp),2)-(dot(temp,temp)-math.Pow(S.radius,2))
  if delta<0 {
    return -1. , 0., 0.
  } else{
    return -dot(R.vel,temp)-math.Sqrt(delta), 0., 0.
  }
}

func (S sphere) get_color( tex texture, p0 float64, p1 float64) color.RGBA {
	return S.color

}
