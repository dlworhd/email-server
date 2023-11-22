package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"
)

type PasswordEmail struct {
	TO          string `json:"to"`
	NEWPASSWORD string `json:"newPassword"`
}

type AuthEmail struct {
	TO     string `json:"to"`
	USERID string `json:"userId"`
}

func main() {
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	username := "dlwodud821@gmail.com"
	password := "eutf hszs dvsa qylp"

	auth := smtp.PlainAuth("", username, password, smtpServer)

	http.HandleFunc("/api/email/join", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			var email AuthEmail

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&email)

			if err != nil {
				http.Error(w, "에러 발생", http.StatusBadRequest)
				return
			}

			to := email.TO
			userId := email.USERID

			sendAuthMail(auth, smtpServer, smtpPort, username, password, to, userId)

		}
	})

	http.HandleFunc("/api/email/password", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			var email PasswordEmail

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&email)

			if err != nil {
				http.Error(w, "에러 발생", http.StatusBadRequest)
				return
			}

			to := email.TO
			newPassword := email.NEWPASSWORD

			sendNewPasswordMail(auth, smtpServer, smtpPort, username, password, to, newPassword)

		}
	})

	fmt.Println("Server is running")
	http.ListenAndServe(":8000", nil)
}

func sendAuthMail(auth smtp.Auth, smtpServer string, smtpPort string, username string, password string, to string, userId string) {

	time := time.Now()
	year, month, day := time.Date()

	subject := "Subject: [ARTX] 회원가입 인증 메일이 도착했습니다.🙌🏻\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<div style=\"width:100%;background-color:#ffffff;margin:0\"><div style=\"width:100%;padding:40px 0;background-color:#ffffff;margin:0 auto;display:block\"><table cellpadding=\"0\" cellspacing=\"0\" border=\"0\" align=\"center\" style=\"margin:0 auto;width:94%;max-width:630px;background:#ffffff;border-width:0;border:0;border-style:solid;box-sizing:border-box\"><tbody><tr style=\"margin:0;padding:0\"><td style=\"width:100%;max-width:630px;margin:0 auto;border-spacing:0;border:0;clear:both;border-collapse:separate;padding:0;overflow:hidden;background:#ffffff\"><div style=\"height:0;max-height:0;border-width:0;border:0;border-color:initial;line-height:0;font-size:0;overflow:hidden;display:none\"></div><div><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" style=\"overflow:hidden;margin:0 auto;padding:0;width:100%;max-width:630px;clear:both;line-height:1.7;border-width:0;font-size:14px;border:0;box-sizing:border-box\" width=\"100%\"><tbody><tr><td><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\"><tbody><tr><td style=\"text-align:center;font-size:0\"><div style=\"max-width:630px;width:100%!important;margin:0;vertical-align:top;border-collapse:collapse;box-sizing:border-box;font-size:unset;display:inline-block\"><div style=\"text-align:left;margin:0;line-height:1.7;word-break:break-word;font-size:16px;font-family:AppleSDGothic,apple sd gothic neo,noto sans korean,noto sans korean regular,noto sans cjk kr,noto sans cjk,nanum gothic,malgun gothic,dotum,arial,helvetica,MS Gothic,sans-serif!important;color:#000000;clear:both;border:0\"><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" style=\"width:100%\"><tbody><tr><td style=\"padding:15px 15px 5px 15px;font-size:16px;line-height:1.7;word-break:break-word;color:#000000;border:0;font-family:AppleSDGothic,apple sd gothic neo,noto sans korean,noto sans korean regular,noto sans cjk kr,noto sans cjk,nanum gothic,malgun gothic,dotum,arial,helvetica,MS Gothic,sans-serif!important;width:100%\"><div><span style=\"font-size:14px\">" + fmt.Sprint(year) + "년" + fmt.Sprint(int(month)) + "월" + fmt.Sprint(day) + "일</span></div><div><span style=\"font-size:20px;font-weight:bold\">드디어..! 인증 메일이 도착했습니다.🙌🏻</span></div><div><span style=\"font-size:12px;color:#8e8f91\">ARTX에서 다양한 영감을 얻어 가시길 바라며,</span></div><div><span style=\"font-size:12px;color:#8e8f91\">인증을 위해 <a href=\"https://ka8d596e67406a.user-app.krampoline.com/api/email/auth?userId=" + userId + "\">인증하기</a>버튼을 클릭해주세요 :)</span></div><div><br></div><div><div><span style=\"font-size:14px\"><img alt=\"🙌\" aria-label=\"🙌\" src=\"https://ci6.googleusercontent.com/proxy/WFy3Hc1Swm4AJ0f0kTFBkuH1iFfhuxqK8sv_Kv1JWhbVyNFfbgVTa16tM7k5QtsbJ6Nf6tSwOaA3mWyB-Lxsc-4g84d0VkXa2UMuiw=s0-d-e1-ft#https://fonts.gstatic.com/s/e/notoemoji/15.0/1f64c/72.png\" class=\"CToWUd\" data-bit=\"iit\"> 회원님:) 오늘도 좋은 아침이에요! 알버트 아인슈타인은 이렇게 말했어요. \"삶은 자전거를 타는 것과 같다. 균형을 유지하기 위해 계속 움직여야 한다.\" 오늘 하루가 균형과 전진을 잃지 않는 하루가 되길 바랄게요</span></div></div></td></tr></tbody></table></div></div></td></tr></tbody></table></td></tr></tbody></table></div></td></tr></tbody></table></div></div>"
	msg := []byte(subject + mime + body)

	err := smtp.SendMail(smtpServer+":"+fmt.Sprint(smtpPort), auth, username, []string{to}, msg)

	if err != nil {
		fmt.Println("이메일 전송 중 에러 발생")
	}
}

func sendNewPasswordMail(auth smtp.Auth, smtpServer string, smtpPort string, username string, password string, to string, newPassword string) {

	time := time.Now()
	year, month, day := time.Date()

	subject := "Subject: [ARTX] 패스워드 초기화 메일이 도착했습니다.🙌🏻\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<div style=\"width:100%;background-color:#ffffff;margin:0\"><div style=\"width:100%;padding:40px 0;background-color:#ffffff;margin:0 auto;display:block\"><table cellpadding=\"0\" cellspacing=\"0\" border=\"0\" align=\"center\" style=\"margin:0 auto;width:94%;max-width:630px;background:#ffffff;border-width:0;border:0;border-style:solid;box-sizing:border-box\"><tbody><tr style=\"margin:0;padding:0\"><td style=\"width:100%;max-width:630px;margin:0 auto;border-spacing:0;border:0;clear:both;border-collapse:separate;padding:0;overflow:hidden;background:#ffffff\"><div style=\"height:0;max-height:0;border-width:0;border:0;border-color:initial;line-height:0;font-size:0;overflow:hidden;display:none\"></div><div><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" style=\"overflow:hidden;margin:0 auto;padding:0;width:100%;max-width:630px;clear:both;line-height:1.7;border-width:0;font-size:14px;border:0;box-sizing:border-box\" width=\"100%\"><tbody><tr><td><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\"><tbody><tr><td style=\"text-align:center;font-size:0\"><div style=\"max-width:630px;width:100%!important;margin:0;vertical-align:top;border-collapse:collapse;box-sizing:border-box;font-size:unset;display:inline-block\"><div style=\"text-align:left;margin:0;line-height:1.7;word-break:break-word;font-size:16px;font-family:AppleSDGothic,apple sd gothic neo,noto sans korean,noto sans korean regular,noto sans cjk kr,noto sans cjk,nanum gothic,malgun gothic,dotum,arial,helvetica,MS Gothic,sans-serif!important;color:#000000;clear:both;border:0\"><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" style=\"width:100%\"><tbody><tr><td style=\"padding:15px 15px 5px 15px;font-size:16px;line-height:1.7;word-break:break-word;color:#000000;border:0;font-family:AppleSDGothic,apple sd gothic neo,noto sans korean,noto sans korean regular,noto sans cjk kr,noto sans cjk,nanum gothic,malgun gothic,dotum,arial,helvetica,MS Gothic,sans-serif!important;width:100%\"><div><span style=\"font-size:14px\">" + fmt.Sprint(year) + "년" + fmt.Sprint(int(month)) + "월" + fmt.Sprint(day) + "일</span></div><div><span style=\"font-size:20px;font-weight:bold\">드디어..! 지구 한 바퀴를 돌아 새로운 패스워드가 도착했습니다.🙌🏻</span></div><div><span style=\"font-size:12px;color:#8e8f91\">ARTX에서 다양한 영감을 얻어 가시길 바라며,</span></div><div><span style=\"font-size:12px;color:#8e8f91\">새로운 패스워드를 알려드립니다. :)</span></div><div><br></div><div><div><span style=\"font-size:14px\">" + newPassword + "</span></div></div></td></tr></tbody></table></div></div></td></tr></tbody></table></td></tr></tbody></table></div></td></tr></tbody></table></div></div>"
	msg := []byte(subject + mime + body)

	err := smtp.SendMail(smtpServer+":"+fmt.Sprint(smtpPort), auth, username, []string{to}, msg)

	if err != nil {
		fmt.Println("이메일 전송 중 에러 발생")
	}
}
