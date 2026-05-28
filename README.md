| Donnée               | Comment c'est représenté en JSON ?                                                            | Comment c'est représenté en XML ?                                                            |
|:---------------------|:----------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------|
| Pays                 | Champ simple `country`                                                                        | Attribut `country` dans `<station>`                                                          |
| Coordonnées          | Objet imbriqué `location` contenant `latitude` et `longitude` dans l'element parent `station` | Balise `<coordinates>` avec attributs `lat` et `lon`dans l'element parent `station`          |
| Altitude             | Champ simple `altitude_m`                                                                     | Attribut `altitude` dans `<coordinates>`                                                     |
| Modèle de capteur    | Objet `device` avec champ `type`                                                              | Attribut `model` dans `<hardware>`                                                           |
| Température          | Champ simple `temperature_celsius` dans chaque observation                                    | Balise `<measure>` avec attribut `type="temperature"` et valeur dans le contenu de la balise |
| Conditions ciel      | Champ simple `conditions`                                                                     | Attribut `sky` dans `<observation>`                                                          |
| Vent                 | Objet imbriqué `wind` avec `speed_kmh` et `direction_deg`                                     | Balise `<wind>` avec attributs `speed` et `direction`                                        |
| Notes (optionnelles) | Champ `notes` pouvant contenir `nul`dans observation                                          | Balise `<note>` présente uniquement si une note existe dans observations                     |

# Weather API — TP J3
API REST météo en Go .

## Lancer le serveur
go run .

Le serveur démarre sur le port `:8080` et charge 30 stations en mémoire au démarrage.
## Routes

| Méthode | Route                       | Description                    | Statuts       |
|---------|-----------------------------|--------------------------------|---------------|
| GET     | /stations                   | Liste toutes les stations      | 200           |
| GET     | /stations/{id}              | Une station précise            | 200, 404      |
| GET     | /stations/{id}/observations | Les observations d'une station | 200, 404      |
| POST    | /stations                   | Crée une station               | 201, 400, 409 |
| PUT     | /stations/{id}              | Remplace ou crée une station   | 200, 201      |
| DELETE  | /stations/{id}              | Supprime une station           | 204, 404      |

## Postman

Testé avec la collection `EFREI_Golang_J3.postman_collection.json` (dossier "Slides 26-27"). Runner : 12/12 OK.