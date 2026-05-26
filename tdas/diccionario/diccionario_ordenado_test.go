package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ABB = []int{500, 1000, 2000, 3000, 5000, 7500}

func compararInt(valorA, valorB int) int {
	if valorA > valorB {
		return 1
	}
	if valorA < valorB {
		return -1
	}
	return 0
}

func TestABBDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestABBDiccionarioClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](compararInt)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestABBUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestABBDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestABBReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestABBDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestABBConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestABBConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	compararBasico := func(valorA, valorB basico) int {
		if strings.Compare(valorA.a, valorB.a) != 0 {
			return strings.Compare(valorA.a, valorB.a)
		}
		return compararInt(valorA.b, valorB.b)
	}

	compararStructs := func(valorA, valorB avanzado) int {
		if compararInt(valorA.w, valorB.w) != 0 {
			return compararInt(valorA.w, valorB.w)
		}
		if strings.Compare(valorA.z, valorB.z) != 0 {
			return strings.Compare(valorA.z, valorB.z)
		}
		if compararBasico(valorA.x, valorB.x) != 0 {
			return compararBasico(valorA.x, valorB.x)
		}
		return compararBasico(valorA.y, valorB.y)
	}

	dic := TDADiccionario.CrearABB[avanzado, int](compararStructs)

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestABBClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestABBValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestABBIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	for i := 0; i < len(cs)-1; i++ {
		require.LessOrEqual(t, strings.Compare(cs[i], cs[i+1]), 0)
	}
}

func TestABBIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestABBIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el diccionario */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioABB(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}

func TestABBIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestABBDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HayAlgoMas())
	primero, _ := iter.VerActual()
	require.EqualValues(t, "Gato", primero)

	iter.Avanzar()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, "Perro", segundo)
	require.True(t, iter.HayAlgoMas())

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, "Vaca", tercero)
	iter.Avanzar()

	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestABBIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Avanzar()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Avanzar()
	segundo, _ := iter3.VerActual()
	iter3.Avanzar()
	tercero, _ := iter3.VerActual()
	iter3.Avanzar()
	require.False(t, iter3.HayAlgoMas())
	require.EqualValues(t, "A", primero)
	require.EqualValues(t, "B", segundo)
	require.EqualValues(t, "C", tercero)
}

func TestABBIteradorConRangoCompleto(t *testing.T) {
	t.Log("Chequea que IteradorRango con desde=nil y hasta=nil se comporte igual que Iterador")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	dic.Guardar("Perro", 1)
	dic.Guardar("Gato", 2)
	dic.Guardar("Vaca", 3)

	iter := dic.IteradorRango(nil, nil)

	require.True(t, iter.HayAlgoMas())
	clave, valor := iter.VerActual()
	require.EqualValues(t, "Gato", clave)
	require.EqualValues(t, 2, valor)
	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, valor = iter.VerActual()
	require.EqualValues(t, "Perro", clave)
	require.EqualValues(t, 1, valor)
	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, valor = iter.VerActual()
	require.EqualValues(t, "Vaca", clave)
	require.EqualValues(t, 3, valor)

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestABBIteradorRangoSinFinalDefinido(t *testing.T) {
	t.Log("Chequea que IteradorRango con desde definido y hasta=nil itere desde la primer clave en el rango")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 1)
	dic.Guardar("B", 2)
	dic.Guardar("C", 3)
	dic.Guardar("D", 4)

	desde := "B"
	iter := dic.IteradorRango(&desde, nil)

	require.True(t, iter.HayAlgoMas())
	clave, _ := iter.VerActual()
	require.EqualValues(t, "B", clave)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, _ = iter.VerActual()
	require.EqualValues(t, "C", clave)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, _ = iter.VerActual()
	require.EqualValues(t, "D", clave)

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestABBIteradorRangoHasta(t *testing.T) {
	t.Log("Chequea que IteradorRango con desde=nil y hasta definido itere elementos hasta llegar al final del rango")

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 1)
	dic.Guardar("B", 2)
	dic.Guardar("C", 3)
	dic.Guardar("D", 4)
	dic.Guardar("E", 5)
	dic.Guardar("F", 6)

	hasta := "C"
	iter := dic.IteradorRango(nil, &hasta)

	require.True(t, iter.HayAlgoMas())
	clave, _ := iter.VerActual()
	require.EqualValues(t, "A", clave)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, _ = iter.VerActual()
	require.EqualValues(t, "B", clave)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, _ = iter.VerActual()
	require.EqualValues(t, "C", clave)

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestABBIteradorRangoDesdeHasta(t *testing.T) {
	t.Log("Chequea que IteradorRango con desde definido y hasta definido itere elementos respetando el rango")

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 1)
	dic.Guardar("B", 2)
	dic.Guardar("C", 3)
	dic.Guardar("D", 4)
	dic.Guardar("E", 5)
	dic.Guardar("F", 6)

	desde := "B"
	hasta := "E"
	iter := dic.IteradorRango(&desde, &hasta)

	require.True(t, iter.HayAlgoMas())
	clave, _ := iter.VerActual()
	require.EqualValues(t, "B", clave)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, _ = iter.VerActual()
	require.EqualValues(t, "C", clave)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, _ = iter.VerActual()
	require.EqualValues(t, "D", clave)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	clave, _ = iter.VerActual()
	require.EqualValues(t, "E", clave)

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestABBIteradorRangoMalDefinido(t *testing.T) {
	t.Log("Chequea que IteradorRango con desde y hasta mal definidos, es decir, sin elementos en el, tire panics")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	dic.Guardar("A", 1)
	dic.Guardar("B", 2)
	dic.Guardar("C", 3)

	desde := "H"
	hasta := "M"
	iter := dic.IteradorRango(&desde, &hasta)

	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestABBIteradorRangoConUnUnicoElemento(t *testing.T) {
	t.Log("Chequea que IteradorRango con desde y hasta iguales devuelve ese unico elemento si pertence al diccionario")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("J", 10)

	desde := "J"
	hasta := "J"
	iter := dic.IteradorRango(&desde, &hasta)

	require.True(t, iter.HayAlgoMas())
	clave, _ := iter.VerActual()
	require.EqualValues(t, "J", clave)
	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestABBIterarRangoCorte(t *testing.T) {
	t.Log("Chequea que IterarRango corte cuando la funcion visitar devuelva false")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	dic.Guardar("A", 1)
	dic.Guardar("B", 2)
	dic.Guardar("C", 3)
	dic.Guardar("D", 4)

	clavesVisitadas := make([]string, 0)
	desde := "B"
	hasta := "D"

	dic.IterarRango(&desde, &hasta, func(clave string, dato int) bool {
		clavesVisitadas = append(clavesVisitadas, clave)
		return clave != "C"
	})

	require.EqualValues(t, []string{"B", "C"}, clavesVisitadas)
}

func ejecutarPruebasVolumenIteradorABB(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el diccionario */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HayAlgoMas())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HayAlgoMas() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Avanzar()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HayAlgoMas(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorABB(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorABB(b, n)
			}
		})
	}
}

func TestABBVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](compararInt)

	/* Inserta 'n' parejas en el diccionario */
	for i := 0; i < 5000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}
