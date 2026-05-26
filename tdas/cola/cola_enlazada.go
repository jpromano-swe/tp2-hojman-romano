package cola

type ColaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(ColaEnlazada[T])
	return cola
}

func crearNodo[T any](elem T) *nodo[T] {
	nuevoNodo := nodo[T]{
		dato:      elem,
		siguiente: nil,
	}
	return &nuevoNodo
}

func (cola *ColaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *ColaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *ColaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := crearNodo(elem)

	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.siguiente = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *ColaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := cola.primero.dato
	cola.primero = cola.primero.siguiente

	if cola.primero == nil {
		cola.ultimo = nil
	}
	return dato
}
