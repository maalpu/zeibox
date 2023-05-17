# zeibox
1. Creamos la Lambda
  - Nombre de la función: zeibox
  - Go 1.x
  - Rol de ejecución
    + (*) Creación de un nuevo rol
  - Configuración avanzada
    + (*) Habilitar URL de la función
      - (*) AWS_IAM
    + (*) Configurar el uso compartido de recursos entre orígenes (CORS)
  - <Crear una función>
  - Configuración del tiempo de ejecución - <Editar>
    + Controlador: main <Guardar>
  - Configuración - Configuración general - <Editar>
    + Memoria: 128 <Guardar>
2. Creamos API Gateway - <Crear API> - API HTTP <Crear>
  - <Agregar integración> Lambda
    + Región: us-east-1
    + Functión de Lambda: zeibox
    + Versión: 2.0
  - Nombre de la API: zeibox <Siguiente>
  - Configurar rutas <Siguiente>
  - Configurar etapas 
    + Nombre de etapa: $default <Eliminar> <Siguiente>
  - <Crear>
  - Routes - (Veremos /zeibox ANY como única ruta creada)
  - CORS <Configurar>
    + Access-Control-Allow-Origin  : *
    + Access-Control-Allow-Headers : *
    + Access-Control-Allow-Methods : *
    + Access-Control-Expose-Headers: *
    + <Guardar>
3. Crear repo en github.com
  - Copiar las carpetas de zeiboxuser
  - Eliminamos: 
    * bd/signup.go
  - Reemplazar: zeiboxuser por zeibox
  - Correr lo siguiente en la línea de comandos: go tidy (Instalará las dependencias)
