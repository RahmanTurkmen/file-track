# File Tracker System

## 1) Qu’est-ce que c’est ?
Projet **File Tracker** composé de :
- **Frontend** : Vue 3 + Pinia (Vite)
- **Backend** : API REST en **Go** (Gin)
- Les fichiers sont stockés dans `backend/data/files.json`.

Le backend simule un traitement :
- `uploading` → `processing` → `done`
- Lors d’un `DELETE`, le statut passe à `deleted` (visible dans l’onglet *Trash* côté frontend).

---

## 2) Structure
- `frontend/` : application Vue
- `backend/` : serveur Go (Gin)
  - `backend/main.go` : routes `/files`
  - `backend/data/files.json` : stockage des fichiers

---

## 3) Backend (Go) — API
Base URL : **http://localhost:8080**

### Routes
- **GET `/files`**
  - Retourne la liste des fichiers (depuis `backend/data/files.json`).

- **POST `/files`**
  - Body JSON attendu :
    ```json
    { "name": "rapport", "size": 500 }
    ```
  - Le serveur renvoie le fichier créé et lance une simulation de traitement.

- **DELETE `/files/:id`**
  - Marque le fichier en `status = "deleted"`.

---

## 4) Frontend (Vue)
Le frontend utilise :
- `frontend/src/services/api.js` (Axios) : `baseURL: http://localhost:8080`
- `frontend/src/stores/files.js` (Pinia store)

### Logique UI
Dans `frontend/src/App.vue` :
- Au chargement : `fetchFiles()`
- Puis rafraîchissement toutes les **2 secondes** (interval)
- Onglets :
  - **My Drive** : `status !== 'deleted'`
  - **Trash** : `status === 'deleted'`
  - **Recent** : `status !== 'deleted'`, tri par `id` décroissant

---

## 5) Commandes pour lancer le projet

### A) Lancer le backend
Dans un terminal :
```sh
cd backend
go run main.go
```
Le backend démarre sur : **http://localhost:8080**.

### B) Lancer le frontend
Dans un autre terminal :
```sh
cd frontend
npm install
npm run dev
```
Le frontend démarre sur une URL Vite (souvent **http://localhost:5173**).

---

## 6) Test rapide
1. Démarrer backend puis frontend.
2. Ouvrir l’interface web.
3. Entrer un **File name** et une **Size MB**, cliquer **Upload**.
4. Observer le statut évoluer (le refresh est automatique toutes les 2s).
5. Cliquer **delete** pour envoyer dans *Trash*.

