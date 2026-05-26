package cola_prioridad

const (
	_CAPACIDAD_INICIAL     = 2
	_FACTOR_DE_CRECIMIENTO = 2
	_FACTOR_DE_REDUCCION   = 4
	_ERROR_HEAP_VACIO      = "La cola esta vacia"
)

type heapDinamico[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heapDinamico[T])
	heap.datos = make([]T, _CAPACIDAD_INICIAL)
	heap.cantidad = 0
	heap.cmp = funcion_cmp
	return heap

}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heapDinamico[T])

	capacidad := len(arreglo)
	if capacidad < _CAPACIDAD_INICIAL {
		capacidad = _CAPACIDAD_INICIAL
	}

	heap.datos = make([]T, capacidad)
	copy(heap.datos, arreglo)
	heap.cantidad = len(arreglo)
	heap.cmp = funcion_cmp
	heap.heapify()

	return heap
}

func (heap *heapDinamico[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heapDinamico[T]) Encolar(elemento T) {
	if heap.cantidad == len(heap.datos) {
		heap.redimensionar(len(heap.datos) * _FACTOR_DE_CRECIMIENTO)
	}
	heap.datos[heap.cantidad] = elemento
	heap.cantidad++
	heap.upheap(heap.cantidad - 1)
}

func (heap *heapDinamico[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(_ERROR_HEAP_VACIO)
	}
	return heap.datos[0]
}
func (heap *heapDinamico[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heapDinamico[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(_ERROR_HEAP_VACIO)
	}
	return heap._desencolar(0, heap.cantidad-1)
}

func (heap *heapDinamico[T]) _desencolar(ini, fin int) T {
	maximoDelHeap := heap.datos[ini]

	if heap.cantidad == 1 {
		heap.cantidad--
		return maximoDelHeap
	}

	heap.datos[ini] = heap.datos[fin]
	heap.cantidad--
	heap.downheap(ini)

	if heap.cantidad > 0 && heap.cantidad*_FACTOR_DE_REDUCCION <= len(heap.datos) {
		heap.redimensionar(len(heap.datos) / _FACTOR_DE_CRECIMIENTO)
	}

	return maximoDelHeap
}

func (heap *heapDinamico[T]) upheap(pos int) {
	if pos == 0 {
		return
	}

	if heap.cmp(heap.datos[pos], heap.datos[(pos-1)/2]) <= 0 {
		return
	}
	if heap.cmp(heap.datos[pos], heap.datos[(pos-1)/2]) > 0 {
		heap.datos[pos], heap.datos[(pos-1)/2] = heap.datos[(pos-1)/2], heap.datos[pos]
	}
	heap.upheap((pos - 1) / 2)
}

func (heap *heapDinamico[T]) downheap(pos int) {
	hijoIzq := (2 * pos) + 1
	hijoDer := (2 * pos) + 2
	aux := pos

	if hijoIzq < heap.cantidad && heap.cmp(heap.datos[hijoIzq], heap.datos[aux]) > 0 {
		aux = hijoIzq
	}
	if hijoDer < heap.cantidad && heap.cmp(heap.datos[hijoDer], heap.datos[aux]) > 0 {
		aux = hijoDer
	}
	if aux == pos {
		return
	}
	heap.datos[pos], heap.datos[aux] = heap.datos[aux], heap.datos[pos]
	heap.downheap(aux)
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapAux := &heapDinamico[T]{
		datos:    elementos,
		cantidad: len(elementos),
		cmp:      funcion_cmp,
	}
	heapAux.heapify()
	heapAux._heapsort()
}

func (heap *heapDinamico[T]) heapify() {
	ultimoPadre := heap.cantidad/2 - 1
	heap._heapify(ultimoPadre)
}

func (heap *heapDinamico[T]) _heapify(pos int) {
	if pos < 0 {
		return
	}
	heap.downheap(pos)
	heap._heapify(pos - 1)
}

func (heap *heapDinamico[T]) _heapsort() {
	if heap.cantidad <= 1 {
		return
	}
	heap.datos[0], heap.datos[heap.cantidad-1] = heap.datos[heap.cantidad-1], heap.datos[0]
	heap.cantidad--
	heap.downheap(0)
	heap._heapsort()
}

func (heap *heapDinamico[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, heap.datos[:heap.cantidad])
	heap.datos = nuevosDatos
}
