package main


import (
	"image/color"
)



type Mat3solve struct{
	M0,M1,M2 vec3
}

func (M Mat3solve) det() float64{
	return M.M0.x*M.M1.y*M.M2.z+M.M1.x*M.M2.y*M.M0.z+M.M2.x*M.M0.y*M.M1.z-M.M2.x*M.M1.y*M.M0.z-M.M0.x*M.M2.y*M.M1.z-M.M1.x*M.M0.y*M.M2.z
}

type triangle struct{
	pos,v0,v1 vec3
	color color.RGBA
	texmap bool
	texpx, texpy int
	texdx, texdy float64
	texori bool
}



func (T triangle) get_color(tex texture,p0 float64, p1 float64) color.RGBA {
	if !T.texmap {
		return T.color
	} else {
		if T.texori {
			return tex.texture[T.texpx+int(T.texdx*p0)][T.texpy+int(T.texdy*p1)]
		} else {
			return tex.texture[T.texpx-1-int(T.texdx*p0)][T.texpy-1-int(T.texdy*p1)]
		}
	}
}


func (T triangle) collide_ray( R ray) (float64,float64,float64) {
	detA:=Mat3solve{R.vel,T.v0,T.v1}.det()
	if detA==0 {
		return -1., 0., 0.
	}else {
		temp:=T.pos.minus(R.pos)
		dist:=Mat3solve{temp,T.v0,T.v1}.det()/detA
		p0:=-Mat3solve{R.vel,temp,T.v1}.det()/detA
		p1:=-Mat3solve{R.vel,T.v0,temp}.det()/detA
		if 0.<p0 && 0.<p1 && p0+p1 <1. {
			return dist,p0,p1
		} else{
			return -1., 0., 0.
		}
	}
}
