// go get -u github.com/go-mail/mail

package email

import (
	"crypto/tls"
	"log"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func loadConfigEmail() {
    viper.SetConfigName("email")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./src/internal/email")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
}

func NewEmailSender(host string, port int, username, password string) *gomail.Dialer {
    d := gomail.NewDialer(host, port, username, password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    return d
}


func EmailSender(from string, to []string, subject, body string) error {
	loadConfigEmail()
    sender := NewEmailSender(viper.GetString("database.host"), viper.GetInt("database.port"), viper.GetString("database.username"), viper.GetString("database.password"))
    message := gomail.NewMessage()
    message.SetHeader("From", from)
    message.SetHeader("To", to...)
    message.SetHeader(subject)
    message.SetBody("text/html", body)

    if err := sender.DialAndSend(message); err != nil {
        log.Fatal("Failed to send email:", err)
    } else {
        log.Println("Email sent successfully!")
    }
	return nil
}