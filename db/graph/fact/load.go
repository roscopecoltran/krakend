package fact

/*

func Load(dbPath string) {
	// We store fact in graph database.
	// Here we use Cayley to store them.
	f := fact.NewFact("")

	// Pass a file name to create a persistent fact
	// f := fact.NewFact(dbPath) // "./db/knowledge.db"

	defer f.Close()

	// Declare truth
	f.Let("แมว").Has("ชื่อ", "เหมียว")
	f.Let("หมา").Has("ชื่อ", "โฮ่ง")
	f.Let("เหมียว").Has("สี", "ดำ")

	fmt.Println(f.What("แมว", "ชื่อ"))
	fmt.Println(f.WhoHas("ชื่อ", "เหมียว"))
	fmt.Println(f.WhoHas("ชื่อ", "โฮ่ง"))
	fmt.Println(f.What(f.What("แมว", "ชื่อ"), "สี"))
	fmt.Println(f.Stringify(f.WhoHas("ชื่อ", "เหมียว")))

	fmt.Println(f.What("เวลา"))
}

*/
