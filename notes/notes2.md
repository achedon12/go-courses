# Suite Types en GO

> slice : un slice est un tableau dynamique qui peut être redimensionné.

# Différentes façons de déclarer une variable :

> var x int = 10 # déclaration explicite avec type
> x := 10 # déclaration avec type inference
> var x = 10 # déclaration avec type inference (type déduit de la valeur)
>
> Possible déclaration multiple :
> var x, y int = 10, 20
> x, y := 10, 20
> var x, y = 10, 20
> /!\ Déclaration multiple avec type différent possible

# Echange de valeur :

> x, y = y, x

# Variable avec plusieurs déclarations groupées

> var (
> x int = 10
> y int = 20
> )


> := est utilisé et déclarable que dans le cadre d'une fonction

# Iota : permet de créer des constantes incrémentales

> const (
> A = iota // A = 0
> B // B = 1
> C // C = 2
> )

# Fonction closure : une fonction qui peut capturer et utiliser des variables de son environnement

=> optimisation de la gestion des ressources

# fallthrough : permet de continuer l'exécution du switch même après avoir trouvé une correspondance

# Tableau avec make : permet de créer un tableau avec une taille dynamique comme un slice

> arr := make([]int, 5) // crée un slice de 5 éléments
> arr := make([]int, 5, 10) // crée un slice de 5 éléments avec une capacité de 10

# append : permet d'ajouter des éléments à un slice

> arr = append(arr, 6) // ajoute l'élément 6 à la fin du slice
> arr = append(arr, 7, 8) // ajoute les éléments 7 et 8 à la fin du slice

# Maps : une collection de paires clé-valeur

> m := make(map[string]int) // crée une map avec des clés de type string et des valeurs de type int
> m["a"] = 1 // ajoute une paire clé-valeur à la map
> value, exists := m["a"] // récupère la valeur associée à la clé "a" et vérifie si elle existe

> Notions de copy des slices et maps : les slices et les maps sont des types de référence, ce qui signifie que lorsque
> qu'on les assigne à une nouvelle variable, vous créez une référence à la même underlying array ou map. Par conséquent,
> les modifications apportées à l'un affecteront l'autre.


kamelCase : convention de nommage où les mots sont collés sans espaces et chaque mot commence par une majuscule, sauf le
premier mot qui commence par une minuscule. Par exemple : myVariableName.
PascalCase : convention de nommage où les mots sont collés sans espaces et chaque mot commence par une majuscule. Par
exemple : MyVariableName.

# Defer : permet de différer l'exécution d'une fonction jusqu'à ce que la fonction englobante retourne. Les fonctions différées sont exécutées dans l'ordre inverse de leur déclaration.

> func main() {
> defer fmt.Println("World")
> fmt.Println("Hello")
>
> // Output:
> // Hello
> // World
> }