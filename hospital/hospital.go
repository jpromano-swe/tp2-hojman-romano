package hospital

import (
  "strconv"
  "strings"
  "tp2/mensajes"
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
  pacientes := diccionario.CrearHash[string, *Paciente]()

  lineas := leerLineasArchivo(archivo)

  for _, linea := range lineas {
    cargarPacienteEnDiccionario(linea, pacientes)
  }
  return pacientes
}

func cargarPacienteEnDiccionario(linea string, pacientes diccionario.Diccionario[string, *Paciente]) {
  datos := strings.Split(linea, ",")
  nombrePaciente := strings.TrimSpace(datos[0])
  anioPaciente := strings.TrimSpace(datos[1])
  antiguedadPaciente, err := strconv.Atoi(anioPaciente)
  if err != nil {
    panic(mensajes.ENOENT_ANIO)
  }
  paciente := crearPaciente(nombrePaciente, antiguedadPaciente)
  pacientes.Guardar(nombrePaciente, paciente)
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
  lineas := leerLineasArchivo(archivo)

  for _, linea := range lineas {
    cargarMedicoEnArbol(linea, arbol, especialidades)
  }
  return arbol
}

func cargarMedicoEnArbol(linea string, arbol diccionario.DiccionarioOrdenado[string, *Medico], especialidades diccionario.Diccionario[string, *Especialidad]) {
  arregloDatos := strings.Split(linea, ",")
  nombreMedico := strings.TrimSpace(arregloDatos[0])
  especialidadMedico := strings.TrimSpace(arregloDatos[1])

  especialidad := obtenerEspecialidad(especialidades, especialidadMedico)
  medico := crearMedico(nombreMedico, especialidad)

  arbol.Guardar(nombreMedico, medico)
}

func obtenerEspecialidad(especialidades diccionario.Diccionario[string, *Especialidad], especialidadMedico string) *Especialidad {

  if especialidades.Pertenece(especialidadMedico) {
    return especialidades.Obtener(especialidadMedico)
  }
  nuevaEspecialidad := crearEspecialidad(especialidadMedico)
  especialidades.Guardar(especialidadMedico, nuevaEspecialidad)
  return nuevaEspecialidad
}
