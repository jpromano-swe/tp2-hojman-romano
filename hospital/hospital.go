package hospital

import (
  "tdas/tdas/cola"
  "tdas/tdas/cola_prioridad"
  "tdas/tdas/diccionario"
)

type Paciente struct {
  nombre     string
  antiguedad int
}

type Medico struct {
  nombre       string
  especialidad *Especialidad
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
