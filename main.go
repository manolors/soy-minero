package main

import (
	"fmt"
	"math/rand"
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

func encontrar(mina Mina, canalMenasEncontradas chan Material) {
	for _, material := range mina {
		if material.tipo == "mena" {
			fmt.Println("Ojeador: encontre una material de", material.material)
			canalMenasEncontradas <- material
		}
	}
	return
}

func minar(canalMenasEncontradas chan Material, canalMenasMinadas chan Material) {
	for {
		m := <-canalMenasEncontradas
		fmt.Println("Minero: minando mena de", m.material)
		canalMenasMinadas <- m
	}
}

func fundir(canalMenasMinadas chan Material) {
	for {
		m := <-canalMenasMinadas
		fmt.Println("Fundidor: fundiendo mina de ", m.material)
		m.tipo = "lingote"
	}
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

func main() {
	var mina Mina
	mina.init()

	canalMenasEncontradas := make(chan Material)
	canalMenasMinadas := make(chan Material)

	go encontrar(mina, canalMenasEncontradas)
	go minar(canalMenasEncontradas, canalMenasMinadas)
	go fundir(canalMenasMinadas)
	<-time.After(time.Second * 5) // ignorar por ahora
}
