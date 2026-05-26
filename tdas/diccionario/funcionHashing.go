package diccionario

import (
	"fmt"
	"hash/fnv"
)

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hasheoDeBytes(arreglo []byte) uint64 {
	hash := fnv.New64a()
	hash.Write(arreglo)
	return hash.Sum64()
}
func asignarIndice(hash uint64, capacidad int) int {
	return int(hash % uint64(capacidad))
}

func hashingDeClaves[K comparable](clave K, capacidad int) int {
	bytes := convertirABytes(clave)
	hash := hasheoDeBytes(bytes)
	return asignarIndice(hash, capacidad)
}
