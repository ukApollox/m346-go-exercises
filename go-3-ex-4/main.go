package main

import "fmt"

const (
    Diamonds = '\u25c6' // Karo (weiße Raute)
    Spades   = '\u2660' // Pik
    Clubs    = '\u2663' // Kreuz
    Hearts   = '\u2665' // Herz

    Six   = '\u2465' // ⑥
    Seven = '\u2466' // ⑦
    Eight = '\u2467' // ⑧
    Nine  = '\u2468' // ⑨
    Ten   = '\u2469' // ⑩
    Jack  = 'J'
    Queen = 'Q'
    King  = 'K'
    Ace   = 'A'
)

func main() {
    suits := []rune{Diamonds, Spades, Clubs, Hearts}
    ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

    // Äußere Schleife: über alle Werte (Zeilen)
    for _, rank := range ranks {
        // Innere Schleife: über alle Farben (Spalten)
        for _, suit := range suits {
            fmt.Printf("%c%c\t", suit, rank)
        }
        fmt.Println() // Zeilenumbruch nach jeder kompletten Reihe
    }
}