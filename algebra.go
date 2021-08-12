package main


import (
	"math"
)

type vec3 struct {
	x,y,z float64
}

func (v0 vec3) minus(v1 vec3) vec3 {
	return vec3{v0.x-v1.x,v0.y-v1.y,v0.z-v1.z}
}

func  dot(v1 vec3, v2 vec3) float64 {
  return v1.x*v2.x + v1.y*v2.y+ v1.z*v2.z
}

func  (v vec3) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y+ v.z*v.z)
}

func (v vec3) normalize() vec3 {
  temp:=v.Abs()
  return vec3{v.x/temp,v.y/temp,v.z/temp}
}

type mat3 struct {
  r0,r1,r2 vec3
}

func mult(M mat3, v vec3) vec3 {
  return vec3{dot(M.r0,v),dot(M.r1,v),dot(M.r2,v)}
}

func T( M mat3) mat3{
  return mat3{vec3{M.r0.x,M.r1.x,M.r2.x},vec3{M.r0.y,M.r1.y,M.r2.y},vec3{M.r0.z,M.r1.z,M.r2.z}}
}

func matmatmult(M0 mat3, M1 mat3) mat3 {
  M1T:=T(M1)
  return mat3{vec3{dot(M0.r0,M1T.r0),dot(M0.r0,M1T.r1),dot(M0.r0,M1T.r2)},
              vec3{dot(M0.r1,M1T.r0),dot(M0.r1,M1T.r1),dot(M0.r1,M1T.r2)},
              vec3{dot(M0.r2,M1T.r0),dot(M0.r2,M1T.r1),dot(M0.r2,M1T.r2)},}
  //return T(mat3{mult(M0,M1T.r0),mult(M0,M1T.r1),mult(M0,M1T.r2)})
}
