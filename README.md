## Repositorio de aprendizaje de [GoLang](https://go.dev/) (19/12/2023)

Tras [instalar Go](https://go.dev/doc/install) y añadir el ejecutable al PATH (comprobar que `go version` devuelve la versión instalada de Go), podemos empezar.

### Primer ejercicio: [Hello World](https://go.dev/doc/tutorial/getting-started)

Por algún sitio hay que empezar... \
Creamos la carpeta [hello](./hello), donde inicializamos un nuevo proyecto básico con la sentencia `go mod init example/hello`, creamos el fichero [hello.go](./hello/hello.go), pegamos el código de ejemplo con un print básico, y ejecutamos con la sentencia `go run hello.go`.\
Todos los programas de Go tienen que estar contenidos en un pacquete, y todos los paquetes pueden tener todos los ficheros .go que queramos, pero sólo puede haber un fichero main.go. Luego desde el main.go podemos invocar a los demás.<br>
Esto es 0 técnico pero qué le vamos a hacer, es un repo de aprendizaje.

### Primer ejercicio.5: Añadiendo un paquete externo como dependencia
Vamos a añadir un paquete externo que genera frases de esas que les flipan a los yanquis.  
Primero vamos a crear una nueva carpeta, [quotes](./quotes), con un fichero en su interior, [quote.go](./quotes/quote.go).  
Por ser nueva carpeta, tenemoq ue volver a ejecutar `go mod init example/quote`.  
Al añadir ahora dependencia de un paquete que no pertenece a la librería básica (¿Go tiene librería básica o cómo se llama?), debemos ejecutar `go mod tidy` para recoger todas las dependencias y, a continuación, podemos ejecutar con `go run quote.go`