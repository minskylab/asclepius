# ASCLEPIUS
Asclepius es un proyecto/plataforma/sistema que pretende servir como tecnológica de entrada para que sistemas de salud tradicionales de países en vías de desarrollo puedan dar un salto digital. Las bases de este proyecto yacen en la necesidad de una modernización correctamente orientada de dichos sistemas de salud, es por esto que Asclepius tiene como principios lo siguiente.

- Modularidad
- Extendibilidad
- Distribución
- Data-driven

Con esos principios en mente, el plan de desarrollo de Asclepius tiene como primer objetivo centrar los esfuerzos de desarrollo en la creación las bases conceptuales y desarrollar un primer módulo para su uso inmediato. 

> El proyecto Asclepius nace en una coyuntura global causada por el virus SARS CoV-2 y con el ánimo de aportar a los diferentes gobiernos del muno se ha llegado a la conclusión de que el primer módulo a ser desarrollado será el Módulo para Emergencias de Salud Pública.  

## Modulos

Como ya se dijo anteriormente, uno de los principios de Asclepius es la creación de software de manera modular que permita un **uso a medida y escalable** de la plataforma, a continuación pasaré a explicar los primeros módulos involucrados en la creación de Asclepius. 

> Estos primeros módulos, a excepción del módulo de entidades, tienen prioridad debido al contexto en el que nace Asclepius, dicho contexto es la emergencia de salud global causada por el virus SARS CoV-2, debido a ello, Asclepius partirá con el módulo para control de emergencias de salud pública para que su implementación apoye en la necesidad de integrar la información de esta emergencia que en la actualidad, Marzo 2020, ya ha afectado a más del 80% del globo.

### Módulo de Entidades

Este módulo esta destinado a gestionar las diferentes entidades dentro del sistema, dichas entidades guardan una relación directa con los tipos de usuarios involucrados en la plataforma.

Las entidades más básicas con las que cuenta dicho módulo son: **Persona**, **Recurso**, **Rol** y **Lugar**. Con dichas entidades, se pueden construir relaciones que permitan describir los diferentes estados de un sistema de salud.

**Persona,** describe a un actor en Asclepius, este puede ser un ciudadano civil, un médico, un gobernante, etc, dichos actores interactuarán entre si en la plataforma. A este nivel, ningún actor tiene diferencias entre sí, todos son iguales y tienen los mismos atributos y funciones. Más adelante se entrará a detallar cómo es que esto funciona.

**Recurso,** un recurso representa un `asset` que están asociados a una persona, algunos ejemplos de recursos son: Un historial médico, una prueba de CoViD-19, un permiso de tránsito, etc. Dichos recursos también tienen en cuenta la asociación de múltiples actores, esto quiere decir que un recurso puede compartir referencias entre muchos, un caso de ejemplo de esto sería: un ciudadano puede ser dueño de un historial médico y al mismo tiempo, dicho historial puede tener notas que tienen a Médicos como autores.

**Rol,** un rol es una entidad que nos sirve para designar funciones a personas en la plataforma, por ejemplo, un grupo de personas puede tener el rol de *ciudadano* y eso les confiere atributos extra que le permiten interactuar de otras maneras con la plataforma, cabe resaltar que los roles no son exclusivos, esto quiere decir que una persona puede tener ilimitados roles y entre estos concederán y/o negarán funcionalidades a dicha persona.

**Lugar, **un lugar, en Asclepius, describe un lugar físico o virtual que sirve para poder contextualizar muchas de las interacciones de las diferentes entidades.



### Módulo de Emergencias de Salud Pública

El módulo de emergencias de salud pública guarda un conjunto de interacciones necesarias para un correcto **seguimiento** de afectados, un **registro del estado local** de la emergencia (más adelante se verá que otro de los principios de Asclepius es la distribución, y que por eso es importante contar con estados locales en este módulo) y finalmente **procedimientos médicos, clínicos y de acción civil** que permitan llevar a cabo la toma de decisiones importantes para sobrellevar la emergencia.

Este módulo de emergencia de salud pública cuenta con 3 importantes componentes.

#### API Público

Este API Público servirá para poder generar soluciones de terceros que necesiten consumir los diferentes recursos solventados por el módulo de emergencias de salud pública. En el futuro se documentará correctamente los protocolos y endpoints disponibles para la "third-party".

#### Clientes Bidireccionales

Los clientes bidireccionales son piezas de software que sirven como punto de anclaje a los usuarios finales del sistema, un ejemplo de cliente es una página web, esta web servirá como punto de entrada para alguno de los actores del módulo de emergencias públicas que les permita interactuar con los diferentes procedimientos involucrados en el ya mencionado módulo. Actualmente existen 2 clientes contemplados:

- **Cliente Aplicativo Móvil**, este cliente no es otra cosa que un aflictivo móvil que permita el rápido acceso, diferenciado*, de los diferentes usuarios finales. Un caso de uso de este aplicativo es la realización de pruebas de tamízale rápido de civid-19 por parte de los ciudadanos, y por otro lado el registro de información que los profesionales de la salud pueden levantar de manera digital para compartirlos de manera remota.
- **Cliente Conversacional,** este cliente no es otra cosa que una agente conversacional (aka chatbot) que es multicanal y esta implementado para Facebook Messenger y WhatsApp**, la finalidad de este cliente es que pueda asistir de manera inmediata y natural a la población en general, no cuenta con las todas las funcionalidades que el cliente Móvil, pero, en contraparte pretende ser intuitivo en su uso, para así, ser accesible para la mayor cantidad de personas.



#### Sistema de Monitoreo

El sistema de monitoreo es el componente encargado de procesar y generar eventos de acción con toda la data recolectada de los diferentes actores de la emergencia que interactuan con los diferentes clientes ya antes mencionados. Este sistema de monitoreo tiene como finalidad generar reportes estadísticamente exactos que serán puestos a disposición a los interesados mediante los APIs públicos y que servirán para una mejor toma de decisiones.

[TODO, Tecnologías propuestas a usar, esquema del sistema en macro y micro y más definiciones para cada concepto poco elaborado]



