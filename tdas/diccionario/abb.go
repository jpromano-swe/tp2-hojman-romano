package diccionario

import TDAPila "tdas/pila"

const (
	_ERROR_ITERADOR           = "El iterador termino de iterar"
	_ERROR_CLAVE_INVALIDA_ABB = "La clave no pertenece al diccionario"
)

type abb[K comparable, V any] struct {
	raiz     *nodo[K, V]
	cantidad int
	cmp      func(K, K) int
}

type nodo[K comparable, V any] struct {
	clave    K
	valor    V
	arbolIzq *nodo[K, V]
	arbolDer *nodo[K, V]
}

type iteradorABB[K comparable, V any] struct {
	pilaNodos  TDAPila.Pila[any]
	inicio     *K
	fin        *K
	comparador func(K, K) int
}

func CrearABB[K comparable, V any](funcCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	nuevoArbol := new(abb[K, V])
	nuevoArbol.cmp = funcCmp
	return nuevoArbol
}

func crearNodo[K comparable, V any](clave K, valor V) *nodo[K, V] {
	nuevoNodo := nodo[K, V]{
		clave:    clave,
		valor:    valor,
		arbolIzq: nil,
		arbolDer: nil,
	}
	return &nuevoNodo
}

func (arbol *abb[K, V]) Obtener(clave K) V {
	nodoConClaveABuscar := arbol.buscarNodo(clave)

	if (nodoConClaveABuscar) == nil {
		panic(_ERROR_CLAVE_INVALIDA_ABB)
	}
	return (*nodoConClaveABuscar).valor
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	return arbol.buscarNodo(clave) != nil
}

func (arbol *abb[K, V]) Guardar(clave K, valor V) {
	posicionDondeGuardar := arbol.buscarLugar(clave)
	if *posicionDondeGuardar != nil {
		(*posicionDondeGuardar).valor = valor
		return
	}
	*posicionDondeGuardar = crearNodo(clave, valor)
	arbol.cantidad++
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	nodoABorrar := arbol.buscarNodo(clave)

	if nodoABorrar == nil {
		panic(_ERROR_CLAVE_INVALIDA_ABB)
	}
	nodoActual := *nodoABorrar
	datoElemento := nodoActual.valor

	if nodoActual.arbolIzq != nil && nodoActual.arbolDer != nil {
		arbol.borrarNodoConDosHijos(nodoABorrar)
	} else if nodoActual.arbolIzq != nil {
		*nodoABorrar = nodoActual.arbolIzq
	} else {
		*nodoABorrar = nodoActual.arbolDer
	}
	arbol.cantidad--
	return datoElemento
}

func (arbol *abb[K, V]) buscarLugar(clave K) **nodo[K, V] {
	return arbol._buscarLugar(&arbol.raiz, clave)
}

func (arbol *abb[K, V]) _buscarLugar(elementoActual **nodo[K, V], clave K) **nodo[K, V] {
	if *elementoActual == nil {
		return elementoActual
	}
	if arbol.cmp(clave, (*elementoActual).clave) == 0 {
		return elementoActual
	}
	if arbol.cmp(clave, (*elementoActual).clave) < 0 {
		return arbol._buscarLugar(&(*elementoActual).arbolIzq, clave)
	}
	return arbol._buscarLugar(&(*elementoActual).arbolDer, clave)
}

func (arbol *abb[K, V]) borrarNodoConDosHijos(nodoABorrar **nodo[K, V]) {
	nodoActual := *nodoABorrar
	nodoAReemplazar := arbol.buscarMaximo(&(nodoActual.arbolIzq))
	nodoReemplazo := *nodoAReemplazar

	nodoActual.clave = nodoReemplazo.clave
	nodoActual.valor = nodoReemplazo.valor

	if nodoReemplazo.arbolIzq != nil {
		*nodoAReemplazar = nodoReemplazo.arbolIzq
	} else {
		*nodoAReemplazar = nil
	}
}

func (arbol *abb[K, V]) buscarMaximo(elemento **nodo[K, V]) **nodo[K, V] {
	if *elemento == nil || (*elemento).arbolDer == nil {
		return elemento
	}
	return arbol.buscarMaximo(&(*elemento).arbolDer)
}

func (arbol *abb[K, V]) buscarNodo(clave K) **nodo[K, V] {
	nodoABuscar := arbol.buscarLugar(clave)
	if *nodoABuscar == nil {
		return nil
	}
	return nodoABuscar
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	arbol._iterarRango(arbol.raiz, desde, hasta, visitar)
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	arbol.IterarRango(nil, nil, visitar)
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

func (iter *iteradorABB[K, V]) HayAlgoMas() bool {
	if iter.pilaNodos.EstaVacia() {
		return false
	}
	if iter.fin == nil {
		return true
	}
	claveActual := iter.pilaNodos.VerTope().(*nodo[K, V]).clave
	return iter.comparador(claveActual, *iter.fin) <= 0
}

func (iter *iteradorABB[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic(_ERROR_ITERADOR)
	}
	nodoActual := iter.pilaNodos.VerTope().(*nodo[K, V])
	iter.pilaNodos.Desapilar()
	if nodoActual.arbolDer != nil {
		iter.apilarRamaIzq(nodoActual.arbolDer)
	}
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic(_ERROR_ITERADOR)
	}
	nodoActual := iter.pilaNodos.VerTope().(*nodo[K, V])
	return nodoActual.clave, nodoActual.valor
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := iteradorABB[K, V]{
		pilaNodos:  TDAPila.CrearPilaDinamica[any](),
		inicio:     desde,
		fin:        hasta,
		comparador: arbol.cmp,
	}
	iter.apilarRango(arbol.raiz)
	return &iter
}

func (iter *iteradorABB[K, V]) apilarRamaIzq(nodoActual *nodo[K, V]) {
	if nodoActual == nil {
		return
	}
	iter.pilaNodos.Apilar(nodoActual)
	iter.apilarRamaIzq(nodoActual.arbolIzq)
}

func (iter *iteradorABB[K, V]) apilarRango(nodoActual *nodo[K, V]) {
	if nodoActual == nil {
		return
	}
	if iter.inicio == nil {
		iter.apilarRamaIzq(nodoActual)
		return
	}
	claveActual := nodoActual.clave
	if iter.comparador(claveActual, *iter.inicio) < 0 {
		iter.apilarRango(nodoActual.arbolDer)
	} else {
		iter.pilaNodos.Apilar(nodoActual)
		iter.apilarRango(nodoActual.arbolIzq)
	}
}

func (arbol *abb[K, V]) _iterarRango(nodoActual *nodo[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodoActual == nil {
		return true
	}

	claveActual := nodoActual.clave
	if desde == nil || arbol.cmp(claveActual, *desde) > 0 {
		if !arbol._iterarRango(nodoActual.arbolIzq, desde, hasta, visitar) {
			return false
		}
	}

	if (desde == nil || arbol.cmp(claveActual, *desde) >= 0) && (hasta == nil || arbol.cmp(claveActual, *hasta) <= 0) {
		if !visitar(claveActual, nodoActual.valor) {
			return false
		}

	}
	if hasta == nil || arbol.cmp(claveActual, *hasta) < 0 {
		if !arbol._iterarRango(nodoActual.arbolDer, desde, hasta, visitar) {
			return false
		}
	}
	return true
}
