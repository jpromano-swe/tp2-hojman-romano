package pila

/* Definición del struct pila proporcionado por la cátedra. */

const _factorDeCrecimiento = 2
const _capacidadInicial = 2
const _factorDeReduccion = 4

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, 0)
	pila.cantidad = 0
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if len(pila.datos) == 0 {
		pila.redimensionar(_capacidadInicial)
	}
	if pila.cantidad == len(pila.datos) {
		pila.redimensionar(len(pila.datos) * _factorDeCrecimiento)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	pila.cantidad--

	dato := pila.datos[pila.cantidad]

	if !pila.EstaVacia() && pila.cantidad*_factorDeReduccion <= len(pila.datos) && len(pila.datos) > _capacidadInicial {
		pila.redimensionar(len(pila.datos) / _factorDeCrecimiento)
	}

	return dato
}

func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, pila.datos[:pila.cantidad])
	pila.datos = nuevosDatos
}
