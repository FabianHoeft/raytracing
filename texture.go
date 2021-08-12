package main


import (
	"image/color"
)

type texture struct{
	texture [][]color.RGBA
}

func gengradient() texture{
	n:=100
	m:=200
	out:=make([][]color.RGBA, m)
	for i := 0; i < m; i++ {
		out[i]=make([]color.RGBA, n)
		for j := 0; j < n; j++ {
			out[i][j]=color.RGBA{0,uint8(200*i/m),uint8(255-100*j/n),0xff }
		}
	}
	return texture{out}
}
