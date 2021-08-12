package main


import (
	"fmt"
	"math"
	"time"
	"image/color"

)




func main() {
	t_start:=time.Now()
	width := 1920
	height := 1080
	cores := 20
	blocksize :=width*height/cores/8

	texture:=gengradient()



  S0:=sphere{vec3{3.,0.,0.},0.5,color.RGBA{20,50,60,0xff}}
  S1:=sphere{vec3{2.,0.,0.},0.125,color.RGBA{1,200,80,0xff}}
	T0:=triangle{vec3{3.,0.,0.},vec3{0.,2.,0.},vec3{0.,0.,1.},color.RGBA{120,20,200,0xff},false,0,0,0.,0.,true}
	T1:=triangle{vec3{3.,0.,0.},vec3{0.,-2.,0.},vec3{0.,0.,1.},color.RGBA{120,20,200,0xff},true,0,0,200.,100.,true}
	T2:=triangle{vec3{3.,-2.,1.},vec3{0.,2.,0.},vec3{0.,0.,-1.},color.RGBA{120,20,200,0xff},true,200,100,200.,100.,false}

	// background:=backsolid{color.RGBA{0,0,0,0}}
	background:=backtexture{getbackground(),1024,256}
  state:=game{90.*math.Pi/180.,height,width,background,cores,blocksize,texture}
  P0:=player{0,math.Pi/2.,vec3{0,0,0}}
  objects:=[]geometry{S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2,S0,S1,T0,T1,T2}
	fmt.Println("objects:",len(objects))


	t_init:=time.Now()
	fmt.Println("init:", t_init.Sub(t_start))
	P0.display(state,objects)
	t_done:=time.Now()

	fmt.Println("render:", t_done.Sub(t_init))


}
