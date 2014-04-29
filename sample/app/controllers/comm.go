package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/iassic/revel-modz/modules/user-files"
	"github.com/iassic/revel-modz/modules/ws_comm"
	"github.com/revel/revel"
	"github.com/revel/revel/mail"
	"os"
	"strconv"
	"strings"
	//"~/gocode/src/github.com/iassic/revel-modz/modules/user-files/userfiles.go"
)

func (c User) Comm(ws *websocket.Conn) revel.Result {
	user := c.userConnected()

	comm := ws_comm.New()
	comm.AddHandler("echo", echoHandler)
	if user != nil {
		comm.AddHandler("filecmd", func(msg string, outbound chan string) {
			uId := user.UserId
			fileCommandHandler(uId, msg, outbound)
		})

		if user.UserId == 200001 {
			comm.AddHandler("email", emailHandler)
		}
	}

	comm.Serve(ws)

	revel.INFO.Println("closing WS connection")
	return nil
}

func echoHandler(msg string, outbound chan string) {
	revel.INFO.Println("Echo: ", msg)
	outbound <- "ack " + msg
}

func emailHandler(msg string, outbound chan string) {
	revel.ERROR.Printf("%q\n", msg)
	if msg == "send test" {
		revel.WARN.Println("sending test message")
		err := sendTestMessage()
		if err != nil {
			outbound <- "error sending message"
			return
		}
		outbound <- "ack email sent"
		return
	}
	outbound <- "error unknown email command"
}

func fileCommandHandler(uId int64, msg string, outbound chan string) { //error (return type)
	var resp_body string
	fields := strings.Fields(msg)
	cmd := fields[0]
	switch cmd {
	case "delete":
		fileIdStr := fields[1]

		//parsedUID := strconv.ParseInt(userIdStr, 0, 64) //may not need
		dsID, err := strconv.ParseInt(fileIdStr, 0, 64)
		if err != nil {
			revel.ERROR.Printf("%s\n", "ERROR Parsing fileIdStr in comm.go")
		}

		revel.WARN.Printf("UID in comm.go fileCommandHandler() for delete: %d\n", uId)
		revel.WARN.Printf("dsID in comm.go fileCommandHandler() for delete: %d\n", dsID)
		deleteOutcome, delErr := userfiles.DeleteUserFileById(TestDB, uId, dsID)
		if delErr != nil {
			revel.WARN.Printf("DELETE FROM DB %t\n", deleteOutcome)
		}
		revel.WARN.Printf("DELETE FROM DB %t\n", deleteOutcome)

		resp_body = "success"

	case "update":

	default:
		resp_body = "Unknown file Command"
	}

	resp := "filecmd_response " + resp_body
	revel.WARN.Printf("%q\n", resp)
	outbound <- resp
	//return nil
}

func sendTestMessage() error {
	mail_server := os.Getenv("MAIL_SERVER")
	mail_sender := os.Getenv("MAIL_SENDER")
	mail_passwd := os.Getenv("MAIL_PASSWD")

	message := mail.NewTextAndHtmlMessage(
		[]string{"demo@domain.com"},
		"Test Message",
		"Test Text Body",
		"Test <b>Html</b> <i>Body</i><br>",
	)
	// message.Cc = []string{"admin@domain.com"}
	// message.Bcc = []string{"secret@domain.com"}
	sender := mail.Sender{
		From:    mail_sender,
		ReplyTo: mail_sender,
	}

	mailer := mail.Mailer{
		Server:   mail_server,
		Port:     587,
		UserName: mail_sender,
		Password: mail_passwd,
		// Host: "iassic.com",
		// Auth: smtp.Auth,
		Sender: &sender,
	}

	return mailer.SendMessage(message)
}
