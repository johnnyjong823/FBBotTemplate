// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"math/rand"
)

var mess = &Messenger{}

func main() {
	port := os.Getenv("PORT")
	log.Println("Server start in port:", port)
	mess.VerifyToken = os.Getenv("TOKEN")
	mess.AccessToken = os.Getenv("TOKEN")
	log.Println("Bot start in token:", mess.VerifyToken)
	mess.MessageReceived = MessageReceived
	
	http.HandleFunc("/webhook", mess.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//MessageReceived :Callback to handle when message received.
func MessageReceived(event Event, opts MessageOpts, msg ReceivedMessage) {
	// log.Println("event:", event, " opt:", opts, " msg:", msg)
	profile, err := mess.GetProfile(opts.Sender.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	answers := []string{"洞么參在此誰敢放肆","就跟你說要有外野了",
						    "洞么參剛到步校","打手槍次數是零",
						    "欸幹嘛不要這樣啊","都欺負我","欸我不是二分隊的啊",
						    "我沒有要簽啊","小心我把你們埋進土裡",
						    "013 洞么參 1800我告訴我家狗我想尿尿它要帶我去，預計1830回家",
						    "冒牌洞么參呢？","我一身哈味","小姐 要不要援交?","看甚麼看 在看把你埋進土裡",
						    "你千萬要相信2000字的心得","洞洞么你的頭怎麼那麼大？"}
	if strings.Contains(msg.Text, "哈哈哈"){
		var txt =answers[rand.Intn(len(answers))] 
		resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("%s", txt))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", resp)
	}else{
		resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Hello , %s %s, 我是哈司令", profile.FirstName, profile.LastName))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", resp)
	}	
}
