import smtplib
from email.mime.text import MIMEText

smtp = smtplib.SMTP('smtp.gmail.com',587)
smtp.ehlo()
smtp.starttls()
smtp.login("linuxware123123@gmail.com","flsnrtm123")

msg = MIMEText('hello')
msg['Subject'] = '테스트'
msg['To'] = 'coke_mania@naver.com'
smtp.sendmail('linuxware123123@gmail.com','coke_mania@naver.com',msg.as_string())

smtp.quit()
