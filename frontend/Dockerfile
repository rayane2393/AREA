# Utilisez l'image officielle de Node.js version 18.17
FROM node:18.17

# Définissez le répertoire de travail dans le conteneur
WORKDIR /app

# Copiez le package.json et le package-lock.json (si disponible) dans le répertoire de travail
COPY package*.json ./

# Installez les dépendances du projet
RUN npm install

# Copiez le reste des fichiers de l'application dans le répertoire de travail
COPY . .

# Exposez le port 8080
EXPOSE 8080

# Définissez la commande par défaut pour exécuter l'application
CMD ["npm", "run", "serve"]
