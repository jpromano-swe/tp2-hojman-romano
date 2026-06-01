package hospital

import (
  "bufio"
  "os"
  "strings"
  "tdas/tdas/cola"
  "tdas/tdas/cola_prioridad"
  "tdas/tdas/diccionario"
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

func crearMedico(nombre string, especialidad *Especialidad) *Medico {
  return &Medico{
    nombre:             nombre,
    especialidad:       especialidad,
    pacientesAtendidos: 0,
  }
}

func crearEspecialidad(nombre string) *Especialidad {
  return &Especialidad{
    nombre:          nombre,
    turnosUrgentes:  nil,
    turnosRegulares: nil,
  }
}

func CrearDiccionarioEspecialidades() diccionario.Diccionario[string, *Especialidad] {
  return diccionario.CrearHash[string, *Especialidad]()
}

func crearArbolMedicos(archivo string, especialidades diccionario.Diccionario[string, *Especialidad]) diccionario.DiccionarioOrdenado[string, *Medico] {
  arbol := diccionario.CrearABB[string, *Medico](strings.Compare)
  aperturaArchivo, _ := os.Open(archivo)
  defer aperturaArchivo.Close()

  scanner := bufio.NewScanner(aperturaArchivo)

  for scanner.Scan() {
    linea := scanner.Text()
    arregloDatos := strings.Split(linea, ",")
    if especialidades.Pertenece(arregloDatos[1]) {
      nuevaEspecialidad := especialidades.Obtener(arregloDatos[1])
      nuevoMedico := crearMedico(arregloDatos[0], nuevaEspecialidad)
      arbol.Guardar(nuevoMedico.nombre, nuevoMedico)
    } else {
      nuevaEspecialidad := crearEspecialidad(arregloDatos[1])
      especialidades.Guardar(nuevaEspecialidad.nombre, nuevaEspecialidad)
      nuevoMedico := crearMedico(arregloDatos[0], nuevaEspecialidad)
      arbol.Guardar(nuevoMedico.nombre, nuevoMedico)
    }
  }
  return arbol
}
