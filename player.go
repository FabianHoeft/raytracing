package main


import (
	"math"
	"image/color"
	"image"
	"os"
	"image/png"
)

type player struct {
  phi,theta float64
  pos vec3
}

func (P player) gen_rays( G game) []ray {

  thetarot:=mat3{vec3{math.Cos(math.Pi/2.-P.theta),0.,-math.Sin(math.Pi/2.-P.theta)},
                vec3{0.,1.,0.},
                vec3{math.Sin(math.Pi/2.-P.theta),0.,math.Cos(math.Pi/2.-P.theta)}}
  phirot:=mat3{vec3{math.Cos(P.phi),-math.Sin(P.phi),0.},
              vec3{math.Sin(P.phi),math.Cos(P.phi),0},
              vec3{0.,0.,1.}}
  totrot:=matmatmult(thetarot,phirot)

  v:=make([]vec3, G.hres*G.vres)
  hinc:=math.Tan(G.Fov/2.)/float64(G.hres)
  vinc:=math.Tan(G.Fov/2.)/float64(G.vres) //*float64(G.hres)/float64(G.vres)
  for i := 0; i < G.vres; i++ {
    for j := 0; j < G.hres; j++ {
      v[i+G.vres*j]=mult(totrot, vec3{1.,vinc*float64(i-G.vres/2),hinc*float64(j-G.hres/2)} ).normalize()
		}
  }

  outrays:=make([]ray, G.hres*G.vres)
  for i := 0; i < G.hres*G.vres; i++ {
    outrays[i]=ray{P.pos,v[i]}
  }
  return outrays
}


func (P player) render( G game, objects []geometry) []color.RGBA {
  output := make([]color.RGBA, G.vres*G.hres)
  for i,rays := range P.gen_rays(G) {
    output[i] = rays.trace( G, objects)
  }
  return output
}

func (P player) display(G game, objects []geometry) {
	var out []color.RGBA
	if G.cores<=1 {
		out=P.render(G,objects)
	}else{
		out=P.render_threaded(G,objects)
	}
	upLeft := image.Point{0, 0}
	lowRight := image.Point{G.vres, G.hres}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x < G.vres; x++ {
			for y := 0; y < G.hres; y++ {
					img.Set(x,y,out[G.vres-1-x+G.vres*(G.hres-1-y)])

			}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}


func (P player) render_threaded( G game, objects []geometry) []color.RGBA {
	output := make([]color.RGBA, G.vres*G.hres)
	rays:=P.gen_rays(G)


	numjobs := len(rays)/G.blocksize
	jobs:= make(chan []ray, numjobs)
	results:= make(chan []color.RGBA, numjobs)
	jobidsend:= make(chan int, numjobs)
	jobidrecive:= make(chan int, numjobs)

	for w := 1; w < G.cores; w++ {
		go trace_thread(w,jobidsend,jobidrecive,jobs, results,objects,G)
	}

	for j := 0; j < numjobs; j++ {
		jobs <- rays[j*G.blocksize:(j+1)*G.blocksize]
		jobidsend<- j
	}

	close(jobs)

	for a := 0; a < numjobs; a++ {
		out:= <-results
		id:=<-jobidrecive
		for i := id*G.blocksize; i < (id+1)*G.blocksize; i++ {
			output[i]=out[i-id*G.blocksize]
		}
	}
	// close(results)
  return output
}

func trace_thread( id int, jobidsend <-chan int, jobidrecive chan<- int, jobs <-chan []ray , results chan<- []color.RGBA,  objects []geometry, G game){
	for raysclice:=range jobs {
		jobid:= <-jobidsend
		output := make([]color.RGBA, len(raysclice) )
		for i,rays:= range raysclice {
			output[i] = rays.trace( G, objects)
		}
		results <- output
		jobidrecive <- jobid
	}
}
