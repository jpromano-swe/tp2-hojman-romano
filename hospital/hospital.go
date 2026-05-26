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

type Doctor struct {
  nombre       string
  especialidad *Especialidad
}

type Especialidad struct {
  nombre             string
  pacientesUrgentes  cola.Cola[*Paciente]
  pacientesRegulares cola_prioridad.ColaPrioridad[*Paciente]
}

type Clinica struct {
  pacientes      diccionario.Diccionario[string, *Paciente]
  especialidades diccionario.Diccionario[string, *Especialidad]
  doctores       diccionario.DiccionarioOrdenado[string, *Doctor]
}
