package main

import(
	"fmt"
)

	var or_arreglo1 = [50] int{}
	var or_arreglo2 = [50] int{}
	var or_arreglo3 = [50] int{}
	var or_arreglo4 = [50] int{}

func mostrar_array(ar[50] int, fin int){

	for i := 0; i < fin; i++{

		fmt.Println(ar[i])
	}
	fmt.Println("____________________________________________")
}

func Burbuja(ListaDesordenada [50]int,fin int) {
 var auxiliar int
 for i := 0; i < fin; i++ {
  for j := 0; j < fin-1; j++ {
   if ListaDesordenada[j] < ListaDesordenada[j+1] {
    auxiliar = ListaDesordenada[j]
    ListaDesordenada[j] = ListaDesordenada[j+1]
    ListaDesordenada[j+1] = auxiliar
   }
  }
 }
}


func main(){

	var numero int
	i := 0
	
	arreglo:= [50] int{}
	arreglo1:= [50] int{}
	arreglo2:= [50] int{}
	arreglo3:= [50] int{}
	arreglo4:= [50] int{}
	

	for{
		fmt.Scanf("%d\n",&numero)
		if(numero == 0){
			break
		}else{
			arreglo[i] = numero
			i++
		}
	}

	division := i/4
	rango := division - 1

	// -------------------- DIVISIONES -------------
	inicio := 0
	contador1 := 0
	contador2 := 0
	contador3 := 0
	contador4 := 0

	for i := 0; i < len(arreglo); i++{

		if(i >= inicio && i <= inicio + rango){
			arreglo1[contador1] = arreglo[i]
			contador1++
		}else if(i >= (inicio + division)*1 && i <= ((inicio + division)*1)+rango){
			arreglo2[contador2] = arreglo[i]
			contador2++
		}else if(i >= (inicio + division)*2 && i <= ((inicio + division)*2)+rango){
			arreglo3[contador3] = arreglo[i]
			contador3++
		}else if(i >= (inicio + division)*3 && i <= ((inicio + division)*3)+rango){
			arreglo4[contador4] = arreglo[i]
			contador4++
		}else{
			break
		}
	}

	// ------------------ MOSTRAR ARRAYS -------------------
	fmt.Println("______________________ARRAYS______________________")
	mostrar_array(arreglo1,division)
	mostrar_array(arreglo2,division)
	mostrar_array(arreglo3,division)
	mostrar_array(arreglo4,division)

	// ------------------- ORDENAMIENTO ----------------

	Burbuja(arreglo1,division)
	Burbuja(arreglo2,division)
	Burbuja(arreglo3,division)
	Burbuja(arreglo4,division)

	// ------------------ MOSTRAR ORDENADOS ---------------

	fmt.Println("---------- ORDENAMIENTO -------------")
	mostrar_array(arreglo1,division)
	mostrar_array(arreglo2,division)
	mostrar_array(arreglo3,division)
	mostrar_array(arreglo4,division)

	fmt.Println("-----------------------")
	fmt.Println(division, rango)	
	
}