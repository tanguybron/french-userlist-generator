french-userlist-generator permet de générer des listes d'utilisateurs français.

Ces listes peuvent être utilisées pour des tests de force brute. Notamment en les utilisant avec Kerbrute en environnement Active Directory afin d'énumérer des utilisateurs sur le domaine.


## Utilisation avec les sources

```bash
go run main.go
```
## Utilisation releases

Des releases sont disponibles pour les systèmes d'exploitation suivants :
- MacOS (pour processeurs arm64)
- Linux (pour processeurs amd64 & arm64)


## Améliorations à suivre

- [ ] Optimisation du code
- [ ] Optimisation des performances (notamment pour la suppression des doublons)

## Crédits

Utilisation d'une librairie pour avoir un terminal interactif : [https://github.com/Nexidian/gocliselect](https://github.com/Nexidian/gocliselect)
