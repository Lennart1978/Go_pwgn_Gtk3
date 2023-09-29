package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/gotk3/gotk3/gtk"
)

const digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!$%&/()=?*~#-_.:,;<>"

var eDigits int

const (
	ui = "<?xml version='1.0' encoding='UTF-8'?> <!-- Generated with glade 3.40.0 --> <interface> <requires lib='gtk+' version='3.24'/><object class='GtkWindow' id='window1'> <property name='can-focus'>False</property> <property name='title' translatable='yes'>Lennart's Password Generator V1.0</property> <property name='resizable'>False</property> <property name='window-position'>center-always</property> <property name='default-width'>200</property> <property name='default-height'>100</property> <child> <!-- n-columns=2 n-rows=3 --> <object class='GtkGrid'> <property name='visible'>True</property> <property name='can-focus'>False</property> <property name='row-homogeneous'>True</property> <property name='column-homogeneous'>True</property> <child> <object class='GtkLabel'> <property name='visible'>True</property> <property name='can-focus'>False</property> <property name='label' translatable='yes'>password:</property> <property name='ellipsize'>start</property> </object> <packing> <property name='left-attach'>0</property> <property name='top-attach'>0</property> </packing> </child> <child> <object class='GtkEntry' id='entryPassword'> <property name='visible'>True</property> <property name='can-focus'>True</property> </object> <packing> <property name='left-attach'>0</property> <property name='top-attach'>1</property> </packing> </child> <child> <object class='GtkButton' id='buttonGenerate'> <property name='label' translatable='yes'>generate</property> <property name='visible'>True</property> <property name='can-focus'>True</property> <property name='receives-default'>True</property> </object> <packing> <property name='left-attach'>0</property> <property name='top-attach'>2</property> </packing> </child> <child> <object class='GtkButton' id='buttonCopy'> <property name='label' translatable='yes'>copy</property> <property name='visible'>True</property> <property name='can-focus'>True</property> <property name='receives-default'>True</property> </object> <packing> <property name='left-attach'>1</property> <property name='top-attach'>2</property> </packing> </child> <child> <object class='GtkLabel'> <property name='visible'>True</property> <property name='can-focus'>False</property> <property name='label' translatable='yes'>digits:</property> </object> <packing> <property name='left-attach'>1</property> <property name='top-attach'>0</property> </packing> </child> <child> <object class='GtkEntry' id='entryDigits'> <property name='visible'>True</property> <property name='can-focus'>True</property> </object> <packing> <property name='left-attach'>1</property> <property name='top-attach'>1</property> </packing> </child> </object> </child> </object> </interface>"
)

func main() {
	// Seed für den Zufallsgenerator initialisieren
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialisiere GTK
	gtk.Init(nil)

	// Erstelle ein neues GTK-Fenster
	builder, err := gtk.BuilderNewFromString(ui)
	if err != nil {
		log.Fatal("Fehler beim Laden der .glade-Datei:", err)
	}

	windowObj, err := builder.GetObject("window1")
	if err != nil {
		log.Fatal("Fehler beim Holen des Fensterobjekts:", err)
	}

	window, ok := windowObj.(*gtk.Window)
	if !ok {
		log.Fatal("Konnte das Fensterobjekt nicht in ein GTK-Fenster umwandeln")
	}

	// Button "buttonGenerate" holen
	buttonGenerateObj, err := builder.GetObject("buttonGenerate")
	if err != nil {
		log.Fatal("Fehler beim Holen des Button-Objekts:", err)
	}

	// Button "buttonCopy" holen
	buttonCopyObj, err := builder.GetObject("buttonCopy")
	if err != nil {
		log.Fatal("Fehler beim Holen des Button-Objekts:", err)
	}

	// entry "entryPassword" holen
	entryPasswordObj, err := builder.GetObject("entryPassword")
	if err != nil {
		log.Fatal("Fehler beim Holen des Entry-Objekts:", err)
	}

	// entry "entryDigits" holen
	entryDigitsObj, err := builder.GetObject("entryDigits")
	if err != nil {
		log.Fatal("Fehler beim Holen des Entry-Objekts:", err)
	}

	// Umwandeln des Entry-Objekts in einen GTK-Entry
	entryDigits, ok := entryDigitsObj.(*gtk.Entry)
	if !ok {
		log.Fatal("Konnte das Entry-Objekt nicht in einen GTK-Entry umwandeln")
	}
	// Wert des Eingabetextes hat sich geändert
	entryDigits.Connect("changed", func() {
		// Umwandeln des Eingabetextes in einen Integer
		str, err := entryDigits.GetText()
		if err != nil {
			log.Println("Fehler beim Holen des Eingabetextes:", err)
		}
		// Umwandeln des Eingabetextes in einen Integer
		eDigits, err = strconv.Atoi(str)
		if err != nil {
			eDigits = 0
			return
		}
	})

	// Umwandeln des Entry-Objekts in einen GTK-Entry
	entryPassword, ok := entryPasswordObj.(*gtk.Entry)
	if !ok {
		log.Fatal("Konnte das Entry-Objekt nicht in einen GTK-Entry umwandeln")
	}

	// Umwandeln des Button-Objekts in einen GTK-Button
	buttonGenerate, ok := buttonGenerateObj.(*gtk.Button)
	if !ok {
		log.Fatal("Konnte das Button-Objekt nicht in einen GTK-Button umwandeln")
	}

	// Umwandeln des copy Buttons in einen GTK-Button
	buttonCopy, ok := buttonCopyObj.(*gtk.Button)
	if !ok {
		log.Fatal("Konnte das Button-Objekt nicht in einen GTK-Button umwandeln")
	}

	// Verknüpfe das "clicked"-Signal des Buttons mit der Funktion "generate"
	buttonGenerate.Connect("clicked", func() {
		entryPassword.SetText(generate())
	})

	// Verknüpfe das "clicked"-Signal des Buttons mit der Funktion "copy"
	buttonCopy.Connect("clicked", func() {
		text, err := entryPassword.GetText()
		if err != nil {
			log.Println("Fehler beim Holen des Eingabetextes:", err)
			return
		}
		// Kopiere den Text in die Zwischenablag
		err = clipboard.WriteAll(text)
		if err != nil {
			log.Println("Fehler beim Kopieren des Passwortes:", err)
			return
		}
	})

	// Schließe das GTK-Fenster, wenn es geschlossen wird
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Zeige das Fenster an
	window.ShowAll()

	// Starte die GTK-Hauptschleife
	gtk.Main()
}

// generate erzeugt ein Passwort mit einer bestimmten Anzahl von Ziffern
func generate() string {
	if eDigits <= 0 {
		return ""
	}
	result := make([]byte, eDigits)
	for i := 0; i < eDigits; i++ {
		// Zufällige Auswahl einer Ziffer aus der "digits"-Zeichenkette
		randomIndex := rand.Intn(len(digits))
		result[i] = digits[randomIndex]
	}
	return string(result)
}
