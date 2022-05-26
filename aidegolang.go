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





// type et struct avec pointeur 

//https://www.golangprograms.com/go-language/struct.html
//Nested Struct Type
