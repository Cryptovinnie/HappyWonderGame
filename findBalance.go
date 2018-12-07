package main

import (
 "fmt"
 "strings"
 "github.com/skycoin/skycoin/src/cipher"
 "os"
 "bufio"
 "math/rand"
 "time"
 "os/exec"
 "bytes"
 "log"
 "encoding/json"
)

var (
 // This is a dev seed and address generated by the Skycoin desktop wallet. There are no funds here.
 testSeed    = "bomb dilemma position use junk flame digital involve ask laugh muscle"// radar"  // add test seed string here
 testAddress = "22rP6UB5qojB9BT8kt9eBAfNQCfkicyg6EU" // add test wallet address here
)

// Balance represents an coin and hours balance
type Balance struct {
	Coins string `json:"coins"`
	Hours string `json:"hours"`
}

// AddressBalances represents an address's balance
type AddressBalances struct {
	Confirmed Balance `json:"confirmed"`
	Spendable Balance `json:"spendable"`
	Expected  Balance `json:"expected"`
	Address   string  `json:"address"`
}


func main() {
	stop := ""
	var loopNumber = 0
	fmt.Println(loopNumber)
	for stop != "1" {
		fmt.Println("------------------------------------")
		var twelveWordSeed [12]string //array with 12 dimensions
		var seedString string
		words := strings.Fields(testSeed)
		fmt.Println(words, len(words))

        wordlength := len(words) //word length from test seed
        fmt.Println("Word count", wordlength)

        //If seed doesn't equal 12 pick remaining words from wordlist.txt
        if wordlength != 12 {
                fmt.Println("Not equal to 12")
                add := 12 - wordlength  //how many words to add
                fmt.Println("Amount of words to add", add)
                // add in words to make seed 12.
                addWordsToSeed := orderSeedWords(loopNumber)//get new words here
                fmt.Println("addSeedWords: ", addWordsToSeed)
				
				//for loop to enter addwords to seed
				for i := 0; i < len(addWordsToSeed); i++{
				words = append(words, addWordsToSeed[i])
				fmt.Println("updatedWords:    ", words)
			}
		
				for i := 0; i < 12; i++ {
				twelveWordSeed[i] = words[i]
				seedString += words[i] + " "
				}
			} else {
				seedString = testSeed
			}

	
 fmt.Println("Entered Seed:    ", testSeed)
 fmt.Println("Seed String:    ", seedString)
 fmt.Println("Address: ", recoverWalletAddressFromSeed(seedString))
	
 jsonResp := string(getAddressBalance(recoverWalletAddressFromSeed(seedString)))

	//fmt.Println("balance: ", jsonResp )
	data := AddressBalances{}
    json.Unmarshal([]byte(jsonResp), &data)
	positiveBalance := data.Confirmed.Coins
	
	fmt.Println(positiveBalance)
	if positiveBalance != "0.000000" {
		fmt.Println("Balance greater then 0:", positiveBalance)
		stop = "1"
	} else {
		fmt.Println("Balance is Zero: ", positiveBalance)
		stop = "0"
		loopNumber ++
	}


}
}

func addSeedWords(n int) []string {
        sum := 0
        var name []string
        random := LinesInFile("Wordlist.txt")
        for sum < n {
        rand.Seed(time.Now().UTC().UnixNano())
        randomnumber := rand.Intn(2047)
        fmt.Println("randomnumber: ", randomnumber)
        name = append(name, random[randomnumber])
        sum ++
}
return name 
}

func orderSeedWords(n int) []string {
	
	var name []string
	random := LinesInFile("Wordlist.txt")
	randomnumber := n
	fmt.Println("Number Seq: ", randomnumber)
	name = append(name, random[randomnumber])

return name 
}

func recoverWalletAddressFromSeed(seed string) string {
	// Generate the first address from the seed
	pk,_,_ := cipher.GenerateDeterministicKeyPair([]byte(seed))
   
	//pk, _ := cipher.GenerateDeterministicKeyPair([]byte(seed))
	addr := cipher.AddressFromPubKey(pk)
	return addr.String()
   }
   
func WordCount(s string) map[string]int {
	   words := strings.Fields(s)
	   m := make(map[string]int)
	   for _, word := range words {
		   m[word] += 1
	   }
	   return m
   }
   
   
func LinesInFile(fileName string) []string {
	   f, _ := os.Open(fileName)
	   // Create new Scanner.
	   scanner := bufio.NewScanner(f)
	   result := []string{}
 // Use Scan.
 for scanner.Scan() {
	line := scanner.Text()
	// Append line to result.
	result = append(result, line)
}
return result
}

func getAddressBalance(address string) string 	{
arg := []string {"addressBalance", address}
fmt.Println(arg)
gopath := os.Getenv("GOPATH")
cmd := exec.Command(gopath +"/bin/skycoin-cli", arg...)
var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	err := cmd.Run()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    outStr, _:= string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	
return outStr
}