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

func encontrar(mina Mina) (m MenasEncontradas) {
	for _, v := range mina {
		if v.tipo == "mena" {
			fmt.Println("Ojeador: encontre una mena de", v.material)
			m = append(m, v)
		}
	}
	return
}

func minar(menas MenasEncontradas) (m MenasMinadas) {
	for i, v := range menas {
		fmt.Println("Minero: miné la mena de la posición", i)
		m = append(m, v)
	}
	return
}

func fundir(menas MenasMinadas) (m Lingotes) {
	for _, v := range menas {
		fmt.Println("Fundidor: fundiendo mina de ", v.material)
		v.tipo = "lingote"
		m = append(m, v)
	}
	return
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
	menasEncontradas := encontrar(mina)
	fmt.Println("Menas encontradas: ", menasEncontradas)
	menasMinadas := minar(menasEncontradas)
	fmt.Println("Menas minadas: ", menasMinadas)
	lingotes := fundir(menasMinadas)
	fmt.Println("Lingotes: ", lingotes)
}
