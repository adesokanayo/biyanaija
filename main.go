package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"os"
)

func main() {
	//PORT := "8080"
	PORT := os.Getenv("PORT")

	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/visits", visits)
	http.HandleFunc("/a", landing)
	http.HandleFunc("/test", test)
	http.HandleFunc("/postdata", postdata)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":"+PORT, nil)

}

func landing(w http.ResponseWriter, r *http.Request) {
	
	var welcomemsg ="*** Welcome to MyPay**** "  + "You can send and receive money from anywhere" 

	fmt.Fprintln(w, welcomemsg)
}

func visits(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		fmt.Println("cookies found ", http.StatusNotFound)
	}
	intvalue := c.Value

	fmt.Fprintln(w, intvalue)

}

func set(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})

	fmt.Fprintln(w, "COOKIE WRITTEN")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Your Cookie:", c)
}

func test(w http.ResponseWriter, req *http.Request) {

	//var a strin

	fs := http.FileServer(http.Dir("assets"))

	tpl, err := template.ParseFiles("assets/index.gohtml")
	http.Handle("/",fs)

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func postdata(w http.ResponseWriter, req *http.Request) {

	sessionid := req.FormValue("sessionid")
	servicecode := 124
	phonenumber := req.FormValue("phonenumber")
	text := req.FormValue("text")
	//var text string

	//var Response string

	fmt.Println("Below are Posted Information")
	fmt.Println("Your SessionID:", sessionid)
	fmt.Println("Your Phone Number:", phonenumber)
	fmt.Println("Your Service code:", servicecode)
	//fmt.Fprintln(w, "Your Text:", text)

	if req.Method == "POST" {
		switch text {
		case "":
			fmt.Fprintln(w, "CON *** Welcome to MyPay Platform ****\n" +  "You can send and receive money easily\n" +
				"1. Request Payment\n"+
				"2. Approve or Reject Pending Payment\n"+
				"Thank you\n")

		case "1":

			fmt.Fprintln(w, "CON Enter Phone number of Payer:")

		case "1*08062224476":

			fmt.Fprintln(w, "CON Enter Amount you want to receive:")
			
		case "1*08062224476*1000":

			fmt.Fprintln(w, "CON Enter Name of Payer")

		case "1*08062224476*1000*mike":

			fmt.Fprintln(w, "Enter your USSD Pin:")
			
		case "1*08062224476*1000*mike*1234":

			fmt.Fprintln(w, "END Mike has been notified to approve the  payment\n" +
				     "You can reach out to Mike\n" +
				"Thank you.\n")

		case "2":

			fmt.Fprintln(w, "CON Below are pending payments requested from you.\n" +
				     "Select the payment you wish to treat:\n" +
				"1. Ayo N3000\n"+
				"2. Jide N2000\n"+
				"3. Shola N500\n"+
				"Thank you\n")

		case "2*1":

			fmt.Fprintln(w, "CON N3000 to be paid to Ayo\n"+
				     "Enter your 4 digit secret code to confirm:"+
				"\n"+

				"Thank you\n")

		case "2*1*1234":

			fmt.Fprintln(w, "END N3000 paid to Ayo\n"+
				"Transaction Successful"+
				"\n"+

				"Thank you\n")

		}

	}

}
