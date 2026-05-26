package diccionario

const (
	_LIBRE estadoDeCelda = iota
	_OCUPADO
	_BORRADO

	_CAPACIDAD_INICIAL     = 7
	_FACTOR_DE_CRECIMIENTO = 2
	_FACTOR_DE_REDUCCION   = 2
	_FACTOR_DE_CARGA       = 0.7
	_CARGA_MINIMA          = 4

	_ERROR_ITERADOR_HASH  = "El iterador termino de iterar"
	_ERROR_CLAVE_INVALIDA = "La clave no pertenece al diccionario"
)

type estadoDeCelda = int

type tablaDeHash[K comparable, V any] struct {
	tablaHash []celdaDiccionario[K, V]
	cantidad  int
}

type celdaDiccionario[K comparable, V any] struct {
	clave         K
	valor         V
	estadoDeCelda int
}
type iteradorHash[K comparable, V any] struct {
	hashPasado   *tablaDeHash[K, V]
	indiceActual int
	posicion     int
}

func (hash *tablaDeHash[K, V]) armadoHash(capacidad int) Diccionario[K, V] {
	hash.tablaHash = make([]celdaDiccionario[K, V], capacidad)
	hash.cantidad = 0
	return hash
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(tablaDeHash[K, V])
	hash.armadoHash(_CAPACIDAD_INICIAL)
	return hash
}

func crearCelda[K comparable, V any](clave K, valor V) celdaDiccionario[K, V] {
	celdaNueva := celdaDiccionario[K, V]{clave: clave, valor: valor, estadoDeCelda: _OCUPADO}
	return celdaNueva
}

func (hash *tablaDeHash[K, V]) Pertenece(clave K) bool {
	_, clavePertenece := hash.buscarIndice(clave)
	return clavePertenece
}

func (hash *tablaDeHash[K, V]) Obtener(clave K) V {
	indiceActual, clavePertenece := hash.buscarIndice(clave)
	if !clavePertenece {
		panic(_ERROR_CLAVE_INVALIDA)
	}
	return hash.tablaHash[indiceActual].valor
}

func (hash *tablaDeHash[K, V]) Borrar(clave K) V {
	indiceActual, clavePertenece := hash.buscarIndice(clave)
	if !clavePertenece {
		panic(_ERROR_CLAVE_INVALIDA)
	}
	valorBorrado := hash.tablaHash[indiceActual].valor
	hash.cantidad--
	hash.tablaHash[indiceActual].estadoDeCelda = _BORRADO
	if hash.cantidad*4 <= len(hash.tablaHash) && len(hash.tablaHash) > _CAPACIDAD_INICIAL {
		hash.redimensionarDiccionario(len(hash.tablaHash) / 2)
	}
	return valorBorrado
}

func (hash *tablaDeHash[K, V]) Guardar(clave K, valor V) {

	indiceActual, clavePertenece := hash.buscarIndice(clave)
	capacidadActual := len(hash.tablaHash)
	factorCargado := float64(hash.cantidad+1) / float64(capacidadActual)

	if clavePertenece {
		hash.tablaHash[indiceActual].valor = valor
		return
	}

	if factorCargado > _FACTOR_DE_CARGA {
		nuevaCapacidad := _FACTOR_DE_CRECIMIENTO * capacidadActual
		hash.redimensionarDiccionario(nuevaCapacidad)
		indiceActual, _ = hash.buscarIndice(clave)
	}
	hash.tablaHash[indiceActual] = crearCelda(clave, valor)
	hash.cantidad++
}

func (hash *tablaDeHash[K, V]) Cantidad() int {
	return hash.cantidad
}

// PRIMITIVAS PARA EL ITERADOR

func (hash *tablaDeHash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	indice := hash.encontrarPrimerOcupado()
	cantidad := 0
	for indice < len(hash.tablaHash) && cantidad < hash.cantidad {
		if hash.tablaHash[indice].estadoDeCelda == _OCUPADO {
			cantidad++
			if !visitar(hash.tablaHash[indice].clave, hash.tablaHash[indice].valor) {
				return
			}
		}
		indice++
	}
}

func (hash *tablaDeHash[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorHash[K, V]{hashPasado: hash, indiceActual: hash.encontrarPrimerOcupado(), posicion: 0}
}

func (iter *iteradorHash[K, V]) HayAlgoMas() bool {
	return iter.posicion < iter.hashPasado.cantidad
}

func (iter *iteradorHash[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic(_ERROR_ITERADOR_HASH)
	}
	return iter.hashPasado.tablaHash[iter.indiceActual].clave, iter.hashPasado.tablaHash[iter.indiceActual].valor
}

func (iter *iteradorHash[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic(_ERROR_ITERADOR_HASH)
	}
	iter.indiceActual = iter.hashPasado.encontrarOcupadoDesdeIndice(iter.indiceActual + 1)
	iter.posicion++
}

func (hash *tablaDeHash[K, V]) buscarIndice(clave K) (int, bool) {
	indiceInicial := hashingDeClaves(clave, len(hash.tablaHash))
	indicePrimerBorrado := -1

	for i := 0; i < len(hash.tablaHash); i++ {
		indiceActual := (indiceInicial + i) % len(hash.tablaHash)
		celdaActual := hash.tablaHash[indiceActual]

		if celdaActual.estadoDeCelda == _OCUPADO && celdaActual.clave == clave {
			return indiceActual, true
		} else if celdaActual.estadoDeCelda == _BORRADO && indicePrimerBorrado == -1 {
			indicePrimerBorrado = indiceActual
		} else if celdaActual.estadoDeCelda == _LIBRE {
			return indiceParaInsertar(indiceActual, indicePrimerBorrado), false
		}
	}
	return indicePrimerBorrado, false
}

func indiceParaInsertar(indiceLibre int, indicePrimerBorrado int) int {
	if indicePrimerBorrado != -1 {
		return indicePrimerBorrado
	}
	return indiceLibre
}

func (hash *tablaDeHash[K, V]) encontrarPrimerOcupado() int {
	return hash.encontrarOcupadoDesdeIndice(0)
}

func (hash *tablaDeHash[K, V]) encontrarOcupadoDesdeIndice(indice int) int {
	for indice < len(hash.tablaHash) && hash.tablaHash[indice].estadoDeCelda != _OCUPADO {
		indice++
	}
	return indice
}

func (hash *tablaDeHash[K, V]) redimensionarDiccionario(nuevaCapacidad int) {
	tablaActual := hash.tablaHash
	cantidadHashViejo := hash.cantidad
	hash.armadoHash(nuevaCapacidad)
	for i := 0; i < len(tablaActual) && hash.cantidad < cantidadHashViejo; i++ {
		if tablaActual[i].estadoDeCelda == _OCUPADO {
			indiceActual, _ := hash.buscarIndice(tablaActual[i].clave)
			hash.tablaHash[indiceActual] = crearCelda(tablaActual[i].clave, tablaActual[i].valor)
			hash.cantidad++
		}
	}
}
