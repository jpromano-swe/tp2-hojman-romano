package pila_test

import (
  "testing"
  TDAPila "tp2/tdas/pila"

  "github.com/stretchr/testify/require"
)

func TestNuevaPilaComienzaVacia(t *testing.T) {
  pila := TDAPila.CrearPilaDinamica[int]()
  require.True(t, pila.EstaVacia())
  require.PanicsWithValue(t, "La pila esta vacia", func() {
    pila.VerTope()
  })
  require.PanicsWithValue(t, "La pila esta vacia", func() {
    pila.Desapilar()
  })
}

func TestPilaVaciadaSeComportaComoVacia(t *testing.T) {
  pila := TDAPila.CrearPilaDinamica[int]()
  pila.Apilar(5)
  require.Equal(t, 5, pila.VerTope(), "Se apilo el numero 5")
  require.False(t, pila.EstaVacia(), "La pila tiene al menos 1 elemento")
  pila.Desapilar()
  require.True(t, pila.EstaVacia(), "La pila esta vacia")
  require.PanicsWithValue(t, "La pila esta vacia", func() {
    pila.Desapilar()
  })
  require.PanicsWithValue(t, "La pila esta vacia", func() {
    pila.VerTope()
  })
}

func TestPilaDeStringsSeComportaComoPila(t *testing.T) {
  pila := TDAPila.CrearPilaDinamica[string]()
  pila.Apilar("Fundamentos")
  pila.Apilar("Estructuras de Datos")
  require.Equal(t, "Estructuras de Datos", pila.VerTope())
  pila.Desapilar()
  require.False(t, pila.EstaVacia(), "La pila tiene al menos 1 elemento")
  pila.Desapilar()
  require.True(t, pila.EstaVacia(), "La pila esta vacia")
}

func TestInvarianteLIFOSeMantiene(t *testing.T) {
  pila := TDAPila.CrearPilaDinamica[int]()
  cantidadElementos := 10
  for i := range cantidadElementos {
    pila.Apilar(i)
    require.False(t, pila.EstaVacia())
    require.Equal(t, i, pila.VerTope())
  }
  require.False(t, pila.EstaVacia())

  for i := cantidadElementos - 1; i >= 0; i-- {
    require.Equal(t, i, pila.VerTope())
    require.Equal(t, i, pila.Desapilar())
  }
  require.True(t, pila.EstaVacia())
}

func TestApilarYDesapilarVolumenNoRompeElPrograma(t *testing.T) {
  pila := TDAPila.CrearPilaDinamica[int]()
  const volumen = 100000

  for i := range volumen {
    pila.Apilar(i)
    require.False(t, pila.EstaVacia())
    require.Equal(t, i, pila.VerTope())
  }
  require.False(t, pila.EstaVacia())

  for i := volumen - 1; i >= 0; i-- {
    require.Equal(t, i, pila.VerTope())
    require.Equal(t, i, pila.Desapilar())
  }
  require.True(t, pila.EstaVacia())
  require.PanicsWithValue(t, "La pila esta vacia", func() {
    pila.Desapilar()
  })
  require.PanicsWithValue(t, "La pila esta vacia", func() {
    pila.VerTope()
  })
}
