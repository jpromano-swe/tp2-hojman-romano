package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNuevaColaComienzaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
}

func TestColaVaciadaSeComportaComoVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(5)
	require.Equal(t, 5, cola.VerPrimero(), "Se encolo el numero 5")
	require.False(t, cola.EstaVacia(), "La cola tiene al menos 1 elemento")
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})
}

func TestColaDeStringsSeComportaComoCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Fundamentos")
	cola.Encolar("Estructuras de Datos")
	require.Equal(t, "Fundamentos", cola.VerPrimero())
	cola.Desencolar()
	require.Equal(t, "Estructuras de Datos", cola.VerPrimero())
	require.False(t, cola.EstaVacia(), "La cola tiene al menos 1 elemento")
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
}

func TestInvarianteFIFOSeMantiene(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cantidadElementos := 10
	for i := 0; i < cantidadElementos; i++ {
		cola.Encolar(i)
	}
	require.False(t, cola.EstaVacia())

	for i := 0; i < cantidadElementos; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
	}
}

func TestEncolarYDesencolarVolumenNoRompeElPrograma(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	const volumen = 100000

	for i := 0; i < volumen; i++ {
		cola.Encolar(i)
	}
	require.False(t, cola.EstaVacia())

	for i := 0; i < volumen; i++ {
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})
}
