package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Material struct {
	tipo     string
	material string
}
type Mina [20]Material

type MenasEncontradas []Material
type MenasMinadas []Material
type Lingotes []Material

func encontrar(mina Mina) {
	for _, material := range mina {
		if material.tipo == "mena" {
			fmt.Println("Ojeador: encontre una material de", material.material)
			canalMenasEncontradas <- material
		}
	}
	close(canalMenasEncontradas)
	wg.Done()
}

func minar() {
	for m := range canalMenasEncontradas {
		fmt.Println("Minero: minando mena de", m.material)
		canalMenasMinadas <- m
	}
	close(canalMenasMinadas)
	wg.Done()
}

func fundir() {
	for m := range canalMenasMinadas {
		fmt.Println("Fundidor: fundiendo mina de ", m.material)
		m.tipo = "lingote"
	}
	wg.Done()
}

func (m *Mina) init() {
	initMateriales := []Material{
		{"roca", "granito"},
		{"roca", "carbon"},
		{"mena", "plata"},
		{"mena", "oro"},
		{"mena", "hierro"},
		{"mena", "cobre"},
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(m); i++ {
		m[i] = initMateriales[rand.Intn(len(initMateriales))]
	}
}

var canalMenasEncontradas, canalMenasMinadas chan Material
var wg sync.WaitGroup

func main() {
	var mina Mina
	mina.init()

	canalMenasEncontradas = make(chan Material)
	canalMenasMinadas = make(chan Material)

	wg.Add(1)
	go encontrar(mina)
	wg.Add(1)
	go minar()
	wg.Add(1)
	go fundir()
	wg.Wait()
}
