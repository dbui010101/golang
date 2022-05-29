// faire une fusion des deux on va dire
https://anupkumarpanwar.medium.com/realtime-chat-app-in-golang-a1e4e9d01fea
https://outcrawl.com/realtime-collaborative-drawing-go





/*birdData := map[string]interface{}{
	"NameRoom": map[string]string{
		"coordonnes": "lllllllllllllllllllllllllllllllllllllllllllllaaaaaaaaaa",
		"admin":  "squak",
		"users": "fjdkslqmj fkdsqj mfjdmslq ",
		"chat": "david: slt\n",
	},
	"total birds": 2,
}
*/

/*
//fmt.Println("va afficher l interieur de salon1", birdData["salon1"].(map[string]string))
// explication ici https://localcoder.org/invalid-operation-type-interface-does-not-support-indexing
*/


// pour ajouter des valeurs dans une map[string]interface{} puis mettre une map[string]string {}

// le premier lien plus facile a comprendre et il y a pas l index 0 alors que le deux oui
https://go.dev/play/p/Nli0eLxxg4


/* premiere etape 
declarer la variable var birdData map[string]interface{} au dessus de la fonction
ensuite dans la fonction
birdData := make(map[string]interface{})

	items := make([]map[string]string, 0)

	item := make(map[string]string)
	item["coordonnes"] = "lllllllllllllllllllllllllllllllllllllllllllllaaaaaaaaaa"
	item["admin"] = "squak"
	item["users"] = "fjdkslqmj fkdsqj mfjdmslq "
	item["chat"] = "david: slt\n"
	items = append(items, item)
	//NameRoom peut etre un string ou une variable qui contient un string ( int aucun idee)
	birdData[NameRoom] = items

	fmt.Println("birdData premier technique",birdData)


	birdData premier technique map["erezr-1653202621878":[map[admin:squak chat:david: slt
 coordonnes:lllllllllllllllllllllllllllllllllllllllllllllaaaaaaaaaa users:fjdkslqmj fkdsqj mfjdmslq ]]]



deuxieme technique
https://go.dev/play/p/z09cAZ3FIu






// pointer pour ne pas pas perdre la valeurs pour int

https://devopssec.fr/article/pointeurs-golang#begin-article-section




//  transformer notre birdata en byte pour renvoyer a html

data, _ := json.Marshal(birdData)
    fmt.Println(string(data))

	return c.ws.WriteMessage(mt, data)


//pointer pour  string 
https://www.golangprograms.com/how-to-declare-and-access-pointer-variable.html
// en dehors de la fonction 

//var nomSalon string ="{"  // variable declaré
//var pointerNomSalon *string    /* pointer variable declaration  devient un pointeur*/
//var NameRoom string


//puis dans la function

//pointerNomSalon = &nomSalon   ensuite va pointer sur la reference nomsalon 

//fmt.Printf("\nAddress of variable: %v", &nomSalon  ) //0blabla adresse de la memoire
//fmt.Printf("\nAddress stored in pointer variable: %v", pointerNomSalon ) //0blanla adresse de la memoire identique a &nomsalon

//fmt.Printf("\nValue of Actual Variable: %s",nomSalon ) // "{"
//mt.Printf("\nValue of Pointer variable: %s",*pointerNomSalon ) // "{" et du avec * il a la meme valeur nomSalon

	NameRoom="david"
	// *pointerNomSalon aura la valeur "{"
	if strings.Contains(*pointerNomSalon, NameRoom) {
		fmt.Println("\nexiste")
		salonExist=1;
	}else{
		fmt.Println("\nexiste pas")
		salonExist=0;
		//fmt.Printf("\ancienne valeur: %s",nomSalon ) 
		//nomSalon=strings.TrimSuffix(nomSalon, "}") // enleve le dernier caractere "}" // pas important
		// nouvelle valeur pour nomSalon
		nomSalon+=NameRoom+"{admin:,users:,coordonnes:,chat:},}"
		//replaceNameRoomAdmin:=NameRoom+"{admin:" pas important
		// ici NameRoom sera une addition de tout les salles
		NameRoom=nomSalon
		//ensuite le pointer va chercher la reference de NameRoom et donc il aura la valeur de NameRoom biensur *pointerNomSalon
		pointerNomSalon= &NameRoom // mega attention ici
		
		//fmt.Printf("\nValue of Actual Variable: %s",nomSalon ) // australia
		fmt.Printf("\nValue of Pointer variable: %s",*pointerNomSalon ) // australia
		
		
	}

// type et truct avec pointeur
// la partie 7
https://devopssec.fr/article/structures-et-methodes-golang#begin-article-section





// type et struct avec pointeur  pas ouf cette technique

//https://www.golangprograms.com/go-language/struct.html
//Nested Struct Type

// best technique 

type AllRoom struct {
	rooms []map[string]string 
}



var notes = make(map[string]string)
notes[nameRoomForUsers] = nameRoomForUsers // ok
notes["strokes"] = ""                      // ok
notes["etape_courante"] = "etape_0"
notes["mot"] = ""
notes[NamePlayer] = "{admin: true, drawer: false, score : 0}"
notes["chats"] = "" // ok
// var notes2 = make(map[string]string)
// notes2["Hatimz"] = "a"
// notes2["Alexz"] = "a"
// notes2["Kevinzzzzz"] = "a"
// notes2 === AllRoom1.rooms[1] cest lindex 1 alors que le truc au dessus  notes cest lindex 0
AllRoom1 = AllRoom{
rooms: []map[string]string{

	/*futurRoom*/ /*nameRoomForUsers*/ notes,
	//notes2, // ajouter a l index 1 cette notes2

},
}



// pour ajouter un nouveau salon de +


var notes = make(map[string]string)
notes[nameRoomForUsers] = nameRoomForUsers
notes["strokes"] = ""
notes["etape_courante"] = "etape_0"
notes[NamePlayer] = "{admin: true, drawer: false, score : 0}"
notes["chats"] = ""
notes["mot"] = ""
AllRoom1.rooms = append(AllRoom1.rooms,

/*futurRoom*/ /*nameRoomForUsers*/
notes,
)



// pour affecter une valeur a la key chat


if salonExist == 1 && strings.Contains(v, "chat") {
//fmt.Println("tableau",listAllData)
//"chat":"david: lu"
//nameRoomForUsers
chat = strings.ReplaceAll(v, "\"chat\":", "")
fmt.Printf("\n  le nouveau message")
//fmt.Println("\n  donnée",AllRoom1 )

// cherche l'index ou se trouve la salle dans Allroom1
var indexAllroom int
for i, item := range AllRoom1.rooms {
	//fmt.Println("index ", v)

	//if ( strings.Contains(item, nameRoomForUsers)){
	//fmt.Println("chaque item", item)
	//fmt.Println("\nitem "+nameRoomForUsers+" :", item[nameRoomForUsers])
	//}
	if item[nameRoomForUsers] == nameRoomForUsers {
		//fmt.Println("nameRoomForUsers ", nameRoomForUsers)
		//fmt.Println("index ", i)
		indexAllroom = i
		// ensuite allons sur allRoom1.rooms[indexAllroom]["chats"]
		//fmt.Println("\n avant chat ", AllRoom1.rooms[indexAllroom]["chats"])
		if strings.Contains(AllRoom1.rooms[indexAllroom]["chats"], chat+"\n") {
			// pour eviter quil ecrit deux fois le nouveau message chat
		} else {
			AllRoom1.rooms[indexAllroom]["chats"] = AllRoom1.rooms[indexAllroom]["chats"] + chat + "\n"
		}

		//fmt.Println("\n apres chat ", AllRoom1.rooms[indexAllroom]["chats"])

	}

}
	
	
// delete user
//{"salon-1":"salon-1","quit":"sam"}
if strings.Contains(string(payload), "quit") {
	deleteSalon := 1
	byt := []byte(string(payload))
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	deleteUser := dat["quit"].(string)
	fmt.Println("\ndeleteuser", deleteUser)
	for i, item := range AllRoom1.rooms {
		//fmt.Println("\ni",i)
		//fmt.Println("\n item ",item)
		if item[nameRoomForUsers] == nameRoomForUsers {

			//fmt.Println(" fait voir",AllRoom1.rooms[i])
			delete(AllRoom1.rooms[i], "\""+deleteUser+"\"")
			//fmt.Println("\n fait voir apres delete user",AllRoom1.rooms[i])

			for _, value := range AllRoom1.rooms[i] {
				if strings.Contains(value, "admin") {
					deleteSalon = 0
				}
			}
			// alors delete salon car il n y a pas plus de user
			if deleteSalon == 1 {
				AllRoom1.rooms = append(AllRoom1.rooms[:i], AllRoom1.rooms[(i+1):]...)
				//delete(AllRoom1,AllRoom1.rooms[i])
				// marche pas???????
				fmt.Println("\n apres delete room", AllRoom1.rooms)
			}
		}

	}
}

	
// choopez les valeurs dans la chaine de string
////{"changement_etape":"etape_2","mot":"outgasses"}
// transforme en  tableau en byte
byt := []byte(string(payload))
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	//fmt.Println(dat)

	changEtape := dat["changement_etape"].(string)
	fmt.Println(" valeur de changement_etape  ", changEtape)

	mot:= fmt.Sprintf("%v", dat["mot"]) // si par exemple la valeur devient bool ou string   chope la valeur par srpintf
