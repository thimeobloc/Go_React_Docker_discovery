# Étape 1 : Utiliser une image de base Golang pour construire l'application
FROM golang:1.19 AS builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers du projet dans le conteneur
COPY . .

# Télécharger les dépendances et construire le binaire
RUN go mod tidy
RUN go build -o simple-api-go .

# Étape 2 : Créer l'image finale
FROM debian:bullseye-slim

# Définir le répertoire de travail dans l'image finale
WORKDIR /app

# Copier le binaire à partir de l'étape précédente
COPY --from=builder /app/simple-api-go .

# Exposer le port 8080
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./simple-api-go"]
