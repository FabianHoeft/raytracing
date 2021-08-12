package main


import (
	"image/color"
)

type ray struct {
  pos, vel vec3
}

func (R ray) trace(G game, objects []geometry) color.RGBA {
  distmin,p0min,p1min:=1e308, 0.,0.
	var index geometry
  for _,S := range objects {
    dist,p0,p1:=S.collide_ray(R)
    if 0<dist && dist<distmin {
      distmin,p0min,p1min=dist,p0,p1
			index=S
    }
  }
  if distmin==1e308 {
    return G.background.color(R.vel)
  } else {
    return index.get_color(G.tex,p0min,p1min)
  }
}
