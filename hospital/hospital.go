package hospital

import (
	"bufio"
	"os"
	"strings"
	"tp2/tdas/cola"
	"tp2/tdas/cola_prioridad"
	"tp2/tdas/diccionario"
)

type Paciente struct {
	nombre     string
	antiguedad int
	urgencia   string
}

type Medico struct {
	nombre             string
	especialidad       *Especialidad
	pacientesAtendidos int
}

type Especialidad struct {
	nombre          string
	turnosUrgentes  cola.Cola[*Paciente]
	turnosRegulares cola_prioridad.ColaPrioridad[*Paciente]
	cantidadEspera  int
}

type Hospital struct {
	pacientes      diccionario.Diccionario[string, *Paciente]
	especialidades diccionario.Diccionario[string, *Especialidad]
	medicos        diccionario.DiccionarioOrdenado[string, *Medico]
}

func crearHospital(archivoDoctores string, archivoPacientes string) *Hospital {
	especialidades := CrearDiccionarioEspecialidades()
	medicos := crearArbolMedicos(archivoDoctores, especialidades)
	pacientes := crearDiccionarioPacientes(archivoPacientes)

	return &Hospital{
		pacientes:      pacientes,
		especialidades: especialidades,
		medicos:        medicos,
	}
}

func CrearHospital(archivoDoctores string, archivoPacientes string) *Hospital {
	return crearHospital(archivoDoctores, archivoPacientes)
}

//Turnos

func crearDiccionarioPacientes(archivo string) diccionario.Diccionario[string, *Paciente] {
	return obtenerDatosCSV(archivo)
}

func compararPacientesPorAntiguedad(PacienteUno, PacienteDos *Paciente) int {
	return PacienteDos.antiguedad - PacienteUno.antiguedad
}

func crearEspecialidad(nombre string) *Especialidad {
	urgentes, regulares := crearColasTurnos(compararPacientesPorAntiguedad)
	return &Especialidad{
		nombre:          nombre,
		turnosUrgentes:  urgentes,
		turnosRegulares: regulares,
		cantidadEspera:  0,
	}
}

func crearPaciente(nombre string, antiguedad int) *Paciente {
	return &Paciente{
		nombre:     nombre,
		antiguedad: antiguedad,
		urgencia:   "",
	}
}

func CrearDiccionarioEspecialidades() diccionario.Diccionario[string, *Especialidad] {
	return diccionario.CrearHash[string, *Especialidad]()
}

// Medico
func crearMedico(nombre string, especialidad *Especialidad) *Medico {
	return &Medico{
		nombre:             nombre,
		especialidad:       especialidad,
		pacientesAtendidos: 0,
	}
}

func crearArbolMedicos(archivo string, especialidades diccionario.Diccionario[string, *Especialidad]) diccionario.DiccionarioOrdenado[string, *Medico] {
	arbol := diccionario.CrearABB[string, *Medico](strings.Compare)
	aperturaArchivo, _ := os.Open(archivo) //REFACTOR: ABORTAR POR ERROR
	defer aperturaArchivo.Close()

	scanner := bufio.NewScanner(aperturaArchivo)

	for scanner.Scan() {
		linea := scanner.Text()
		arregloDatos := strings.Split(linea, ",")
		nombreMedico := strings.TrimSpace(arregloDatos[0])
		especialidadMedico := strings.TrimSpace(arregloDatos[1])
		var nuevaEspecialidad *Especialidad //REFACTOR

		if especialidades.Pertenece(especialidadMedico) {
			nuevaEspecialidad = especialidades.Obtener(especialidadMedico)
		} else {
			nuevaEspecialidad = crearEspecialidad(especialidadMedico)
			especialidades.Guardar(especialidadMedico, nuevaEspecialidad)
		}
		nuevoMedico := crearMedico(nombreMedico, nuevaEspecialidad)
		arbol.Guardar(nuevoMedico.nombre, nuevoMedico)
	}
	return arbol
}

//cargarMedicoEnArbol

//crearObtenerEspecialidad

//leerLineasArchivo
