# Hopfield Model

En este repositorio se puede encontrar una implementación para el modelo 
de red neuronal de Hopfield en Go. Esta implementación entrena el modelo
usando el método de Hebb y muestra resultados favorables probando
el modelo con un 20% de ruido aleatorio.

## Estructura del proyecto

```
// Datos de entrenamiento
data/
-- centered.txt
-- uncentered.txt
// Contiene la implementación de la red Hopfield
models/
-- hopfield.go
// Funciones y utilidades para
// * Cargar matrices del sistema de archivos
// * Mostrarlas por pantallla
utils/
-- files.go
-- matrix.go
// Programa principal
cmd/
-- poc/
---- main.go
```

## Implementación

El modelo se implementa mediante una estructura de datos `HopfieldModel`,
el cual se puede instanciar sin argumentos con `NewHopfieldModel` y entrenar con el método `TrainWithHebbMethod`.
El número de nodos se deduce durante la fase de entrenamiento.
Cabe destacar que este modelo lee y retorna matrices NxM, es decir, imagenes
de N pixeles de alto por M píxeles de ancho, sin embargo, internamente realiza
las transformaciones necesarias para poder trabajar con vectores NM, es decir
matrices NMx1.

```plantuml
@startuml
class HopfieldModel {
    TrainWithHebbMethod(patterns Dense[])
    PrintWeights()
    PrintDims()
    Pass(input Dense)
}
@enduml
```

## Resultados

### Primer patron

Corresponde a una imagen con un círculo en el centro.

```
0 0 0 0 0 0 0 0 0 0 
0 0 0 0 0 0 0 0 0 0 
0 0 0 0 0 0 0 0 0 0 
0 0 0 0 1 1 0 0 0 0 
0 0 0 1 0 0 1 0 0 0 
0 0 0 1 0 0 1 0 0 0 
0 0 0 0 1 1 0 0 0 0 
0 0 0 0 0 0 0 0 0 0 
0 0 0 0 0 0 0 0 0 0 
0 0 0 0 0 0 0 0 0 0
```

Agregando 20% de ruido, la imagen se ve de la siguiente manera

```
0 0 0 0 0 0 1 0 1 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 1 0 0 1 0 0 1 0 0
0 0 0 1 0 0 1 0 0 0
0 0 0 0 0 1 1 0 0 0
0 0 0 0 1 1 0 0 0 0
1 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 1 0
0 0 0 0 0 0 0 0 0 0
```

Al hacer la pasada por el modelo, se obtiene la imagen esperada, el primer patron.

```
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 1 1 0 0 0 0
0 0 0 1 0 0 1 0 0 0
0 0 0 1 0 0 1 0 0 0
0 0 0 0 1 1 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
```

### Estadísticas

Si realizamos 100 experimentos de Bernoulli para conocer el rendimiento de la red neuronal,
en particular, el porcentaje de veces que infiere la clase correcta, obtenemos los siguientes resultados.

```
20% Noise. Stabilization with 1 passes. Success ratio of 1.00
30% Noise. Stabilization with 91 passes. Success ratio of 0.90
40% Noise. Stabilization with 161 passes. Success ratio of 0.67
```

## Licencia

Copyright 2022 Carlos David Gonzalez Nexans

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
