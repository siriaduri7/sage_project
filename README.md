# sage_project (jul1 hw)

slices:
when you have list of numbers than we use slices.

Raj sir classes 
This is repo to maintain notes in the golang
// in functions if there 2files under same folder so those functions can accessed in all files under same package like main thats  why i can able to access deck.go and main.go in deck of cards

if you have variable an dif you dont wnat to use than we give "_" eg: for i, cards := range d { // so here i dnt need i so i can keep _ // for _, cards := range d {}

}
#users ApI

- Create User
- Read / List User
- Update User
- Delete User

#packages

- gorilla/mux

##### ASSIGNMENT 1

- Add a controller to accept POST request and print your name in console.

http://localhost:8080/ping  //when ever somecall this port it always goes here (sample.http)

func stratapplication() {
    C:="srika"

} // it is consider to be open local scope even it is upper case it cannot be used outside of the scope, only global can use outside so check it always 

HTTp is protocol to communicate over i/n thats y u see http in browser , so it communiacte using http 


// go mod commands //
go mod init
go list -m all
cat go.mod

 /////////////////HW/////////////////////
 packages(2packges, 3 multiple function, 1 unexported) main lo call chesi execute chesi chupinchali

 2)  create interface shud have (3 methods, 2 structs interfaces ni implement cheyali, chesi oka main lopala a = strcut1 b= struct 2 a function ki a pass chestey a execute avali fn ki b pass chestey b execute avali, print avali, ///////////////my task./////////////

func main() {
newdeck := []string{"Ace of spade","two of the daimonds}
newdeck = append{newdeck, "three of clubs"}

printdeck(newdeck)

func printdeck(d []string) {
    fmt.Println(d)
}

/////type deck []string /// instead of using []string(slice of string everywer) we can use deck instead of []string