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
		case "*384*9294*23312345678*1000#":
			fmt.Fprintln(w, "CON *** Welcome to MyPay****\n" +  "Select Debit Account\n" +
				"1. 00*****211\n"+
				"2. 06*****790\n")
				

		case "1":

			fmt.Fprintln(w, "CON Enter Receiver Account Number:")

		case "1*0000000033":

			fmt.Fprintln(w, "CON Select Payout Currrency:\n" +
			
				"1. Ghanian Cedis (GHS)\n"+
				"2. United States DOllar (USD)\n")
		
			
		case "1*0000000033*1":

			fmt.Fprintln(w, "CON Enter Payment Purpose:")

		case "1*0000000033*1*School Fees":

			fmt.Fprintln(w, "CON You want to send GHS 1,000 to 0000000033. Fee NGN190.00 Total Debit: NGN15,000.00\n" +
		"Enter Pin: \n")
			
		case "1*0000000033*1*School Fees*1234":

			fmt.Fprintln(w, "END Your Payment has been received and will be processed in 24 hours.\n" +
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
