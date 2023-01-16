package main

import (
	"fmt"
	"unitesttraining/src/api/providers/locations_provider"
)

type Persona struct {
	name     string
	lastName string
	age      int
}

func main() {
	country, err := locations_provider.GetCountry("AR")

	fmt.Println(err)
	fmt.Println(country)
	/*
		token := "Hola buenas tardes" + time.Now().UTC().String()
		cypher := sha256.New()

		cypher.Write([]byte(token))

		resToken := string(cypher.Sum(nil))
		sToken := hex.EncodeToString([]byte(resToken))
		fmt.Println(sToken)
		fmt.Println(evaluatePersona("Miguel"))
	*/
}

func evaluatePersona(str string) Persona {
	m := make(map[string][]string)

	m["Hola"] = []string{"1", "2", "3", "4", "5", "6", "7"}
	fmt.Println(m)

	p := Persona{}
	// switch case example
	name := str
	switch name {
	case "Miguel":
		fmt.Println("Is miguel")
		p = Persona{
			name:     "Miguel",
			lastName: "Rodriguez",
			age:      30,
		}
	case "Antonio":
		fmt.Println("Is Antonio")
		p = Persona{
			name:     "Antonio",
			lastName: "Rodriguez",
			age:      31,
		}
	case "Gael":
		fmt.Println("Is Gael")
		p = Persona{
			name:     "Gael",
			lastName: "Rdz",
			age:      5,
		}
	case "Abril":
		fmt.Println("Is Abril")
		p = Persona{
			name:     "Abril",
			lastName: "Rdz",
			age:      3,
		}
	}
	return p
}
