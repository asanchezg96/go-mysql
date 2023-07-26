# CRUD GO MYSQL

Este es un proyecto basico donde se crear un crud usando go y mysql, este crud contiene un modelo user que contiene la logica de negocios.
Ademas contiene un script en SQL para creacion de la base de datos en mysql junto a su respectivo usuario.
Es importante mencionar que crea un schema desde el lenguaje go hacia la base datos MYSQL.


## Instalacion y dependencia
 1. Asegurate de tener instalado go
 2. Dirigete al directorio principal, y ejecuta en la linea de comando: 
  - go mod init go-mysql-crud                  "Modulariza el programa"
  - go get -u github.com/go-sql-driver/mysql   "Agrega driver mysql para generar la coneccion"

**Nota:** Recuerda instalar estas dependencias antes de ejecutar el proyecto.