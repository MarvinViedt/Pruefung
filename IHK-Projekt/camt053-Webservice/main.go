/*
##############################################################################################################
	Projekt: CAMT053-Webservice
	Autor: Marvin Viedt
	Copyright: a.b.s. Rechenzentrum GmbH
	Version: 1.0 (still WIP)
	Release: Sommer 2020
##############################################################################################################
*/

package main

import (
	"fmt"
	"log"
	"errors"
	"net/http"
	"html/template"
	"time"
	"strings"
	"regexp"
	"os"
	"github.com/antchfx/xquery/xml"
	"path/filepath"
)

var ( 
	Map = make(map[string]string)
)

func main() {

	log.Println("Server wird gestartet")

	http.HandleFunc("/", handler)
	http.HandleFunc("/Camt053/", kauszughandler)
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	err := (http.ListenAndServe(":3000", nil))

	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {

	t, err := template.New("Webservice").Parse(webseite)
	fmt.Println("CAMT053-Webservice wird ausgeführt")

    expiration := time.Now().Add(10 * time.Second)
    cookie := http.Cookie{Name: "Session", Value: "Session", Expires: expiration}
    http.SetCookie(w, &cookie)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    err = t.Execute(w, nil)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func kauszughandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Das ist nicht gültig", http.StatusMethodNotAllowed)
		return
	}

	formvalue, err := getformvalue(r)

	if err != nil {
		http.Error(w, fmt.Sprintf("Formeingabe(n) konnte(n) nicht geladen oder bearbeitet werden: %s", err), http.StatusBadRequest)
		return
    }

	if r.URL.Path == "/Camt053/Kontoauszugs-View" {
		var kontoauszug = true
		fmt.Println("Test 1")
		err = readxmlfile(kontoauszug, formvalue, r, w)
		fmt.Println("\nTest 3")
        if err != nil {
			http.Error(w, "Einlesen der XML Datei ist fehlgeschlagen", http.StatusBadRequest)
			return
        }
	}
	
	if r.URL.Path == "/Camt053/XML-View" {
		var kontoauszug = false
		fmt.Println("Test 1")
		err = readxmlfile(kontoauszug, formvalue, r, w)
		fmt.Println("\nTest 3")
        if err != nil {
			http.Error(w, "Einlesen der XML Datei ist fehlgeschlagen", http.StatusBadRequest)
			return
        }
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func getformvalue(r *http.Request) (string, error) {
	err := r.ParseForm()

	if err != nil {
		log.Println("Error:", err)
	}

	formvalue := r.FormValue("XMLFile")
        fmt.Println(formvalue)

	if err != nil {
		errors.New(fmt.Sprintf("Parsen des Eingabefeldes XMLFile fehlgeschlagen: %s", err))
	}

	return formvalue, nil
}

func checkstring(s string) bool {

	if (s >= "A" && s <= "Z") || (s >= "0" && s <= "9") || (s == "\n") {
		return true
	}

	return false
}

func gettagname(get_tag string, w http.ResponseWriter) string {

	value := regexp.MustCompile("((<.*?>))")

	tags := value.FindString(get_tag)
	rmrightbracket := strings.ReplaceAll(tags, ">", "")
	tagname := strings.ReplaceAll(rmrightbracket, "<", "")

	return tagname
}

func getxmlvalues(xmlcontent *xmlquery.Node, r *http.Request, w http.ResponseWriter) [200]string {

	var i int
	var xmlarr [200]string
    for ; xmlcontent != nil; xmlcontent = xmlcontent.NextSibling {

		get_xmlcontent := xmlcontent.InnerText()
		get_xmltag := xmlcontent.OutputXML(true)

		trim_xmlcontent := strings.Trim(get_xmlcontent, "\t \n")
		rmnewlines := regexp.MustCompile("\n\n\n")
        content_rmnewlines := rmnewlines.ReplaceAllString(trim_xmlcontent, "")
        rmtab := regexp.MustCompile("\t")
        content_rmtab := rmtab.ReplaceAllString(content_rmnewlines, "")
        rmspaces2 := regexp.MustCompile("\n\n")
        content_rmspaces2 := rmspaces2.ReplaceAllString(content_rmtab, "\n")
        rmspaces3 := regexp.MustCompile("\n\n")
        content_rmspaces3 := rmspaces3.ReplaceAllString(content_rmspaces2,"\n")

		check := checkstring(content_rmspaces3)

        if check == true  {

        //An dieser Stelle alle Elemente und vllt der Tag Ausgabe soll die Zuweisung sein.
					
		tagname := gettagname(get_xmltag, w)
		tagvalue := content_rmspaces3
					
		Map["tagname"] = tagname

		//fmt.Fprintf(w, "%s\n", tagname)
		//fmt.Fprintf(w, "%s\n", tagvalue)
		
		
		xmlarr[i] = tagvalue
		i++

        } 
    }
	
	//fmt.Fprintf(w, "%s\n", xmlarr)
	
	return xmlarr
}

func readxmlfile(check bool, filename string, r *http.Request, w http.ResponseWriter) error {

	fmt.Println("Test 2\n")

	pwd, err := filepath.Abs("./XML/" + filename + ".xml")
	if err != nil {
		fmt.Println("Die XML Datei konnte nicht gefunden werden: %s", err)
		return err
	}
	
	fmt.Println("Der Dateipfad ist: %s\n", pwd)
	getxmlcontent, _ := os.Open(pwd)

	parsexmlcontent, err := xmlquery.Parse(getxmlcontent)
	if err != nil {
		fmt.Println("Öffnen der XML Datei ist fehlgeschlagen: %s", err)
		return err
	}

	fmt.Println("-----------------------------------------------------------------")

	grphdr := xmlquery.FindOne(parsexmlcontent, "//BkToCstmrStmt/GrpHdr/MsgId")
	grphdrarr := getxmlvalues(grphdr, r, w)

	stmt := xmlquery.FindOne(parsexmlcontent, "//BkToCstmrStmt/Stmt/Id")
	stmtarr := getxmlvalues(stmt, r, w)
	
	//fmt.Fprintf(w, "%s\n", Map)	
	fmt.Println("%s\n", stmtarr)

	if check == true {
		printauszugcontent(w, grphdrarr, stmtarr) 
	}
	if check == false {
		printxmlcontent(w, grphdrarr, stmtarr) 
	}

	return nil
}

func printauszugcontent(w http.ResponseWriter,grphdrarr [200]string, stmtarr [200]string) {
	t := template.Must(template.New("Kontoauszugs-View").Parse(KontoauszugsView))
	
	Nm := strings.Fields(grphdrarr[2])
	IBAN := strings.Fields(stmtarr[4])
	stmt1 := strings.Fields(stmtarr[6])
	
	stmt2 := strings.Fields(stmtarr[8])
	
	//fmt.Fprintf(w, "%s\n", IBAN)
	
	Map["MsgId"] = grphdrarr[0]
	Map["CreDtTm"] = grphdrarr[1]
	Map["Nm"] = Nm[0] + " " + Nm[1] + " " + Nm[2] + " " + Nm[3]
	
	Map["IBAN"] = IBAN[0]

	CreDtTm := grphdrarr[1]
	
	Map["CreateDate"] = getdate(CreDtTm)
	Map["buchdt"] = getdate(stmt2[3])
	Map["valdt"] = getdate(stmt2[4])
	
	kntnr := IBAN[0]
	Map["Kntnr"] = string(kntnr[len(kntnr)-7:])
	
	blz := strings.Fields(strings.Trim(IBAN[0], "DE25 7404000"))
	Map["BLZ"] = blz[0]
	
	Map["BIC"] = IBAN[12]
	Map["Kntst"] = stmt1[1]
	

	Map["ums"] = stmt2[0]
	
	Map["zweck"] = strings.ToUpper(stmt2[len(stmt2)-1])
	Map["zwecknm"] += stmt2[41] + " " + stmt2[42] + " " + stmt2[43]
	for i := 47; i < len(stmt2)-1; i++ {
		Map["info"] += " " + stmt2[i]
	}
	Map["kndref"] = stmt2[11]
	Map["manref"] = stmt2[13]
	
	err := t.Execute(w, Map)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	Map = make(map[string]string)
}

func printxmlcontent(w http.ResponseWriter,grphdrarr [200]string, stmtarr [200]string) {
	t := template.Must(template.New("XML-View").Parse(XMLView))
	
	var grphdrcontent string
	var stmtcontent string
	var newgrphdrarr []string
	var newstmtarr []string

	newgrphdrarr = remove_empty_value (grphdrarr)
	
	for i:=0; i < len(newgrphdrarr); i++ {
		grphdrcontent += "\n" + newgrphdrarr[i]
	}
	
	newstmtarr = remove_empty_value (stmtarr)
	
	for i:=0; i < len(newstmtarr); i++ {
		stmtcontent += "\n" + newstmtarr[i]
	}
	
	Map["grphdr"] = grphdrcontent
	Map["stmt"] = stmtcontent
	
	err := t.Execute(w, Map)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func remove_empty_value (arr [200]string) []string {
    var newarr []string
    for _, str := range arr {
        if str != "" {
            newarr = append(newarr, str)
        }
    }
    return newarr
}

func getdate(datearr string) string {

	getdate := string(datearr[0:10]) 
	layout := "2006-01-02"
	resdate, _ := time.Parse(layout, getdate)
	dateString := resdate.Format("2.01.2006");
	
	return dateString
}



























