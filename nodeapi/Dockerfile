# Usar una imagen base de Node.js
FROM node:14-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de configuración de paquetes al directorio de trabajo
COPY package*.json ./

# Instalar las dependencias del proyecto
RUN npm install

# Copiar todos los archivos del proyecto al contenedor
COPY . .

# Exponer el puerto en el que tu aplicación estará escuchando
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["node", "app.js"]
