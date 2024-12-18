import imaplib
import email
import requests
from email.header import decode_header
import time
from getpass import getpass

IMAP_SERVER = "mail.abdtyx.click"
EMAIL = "admin@abdtyx.click"
PASSWORD = getpass()
OPTIMAILAPI = "http://localhost:80/api"

with imaplib.IMAP4(IMAP_SERVER) as mail:
    mail.login(EMAIL, PASSWORD)
    while 1:
        mail.select("inbox")

        status, messages = mail.search(None, "ALL")
        for num in messages[0].split():
            status, data = mail.fetch(num, "(RFC822)")
            if status != "OK":
                print("Failed to fetch email.")
                continue
            raw_email = data[0][1]
            msg = email.message_from_bytes(raw_email)
            # get sender
            from_header = msg["To"]
            sender_email = None
            if from_header:
                decoded_header = decode_header(from_header)
                for part, encoding in decoded_header:
                    if isinstance(part, bytes):
                        part = part.decode(encoding or "utf-8")
                    if "<" in part and ">" in part:
                        sender_email = part.split("<")[1].strip(">")
                        break
                    elif "@" in part:
                        sender_email = part
                        break
            if sender_email == EMAIL:
                continue
            # get body
            body = ""
            emailaddr = msg['']
            if msg.is_multipart():
                for part in msg.walk():
                    content_type = part.get_content_type()
                    content_disposition = str(part.get("Content-Disposition"))
                    if content_type == "text/plain" and "attachment" not in content_disposition:
                        body += part.get_payload(decode=True).decode()
                        # print("Body:", body)
            else:
                body += msg.get_payload(decode=True).decode()
                # print("Body:", body)
            # call Optimail api
            r = requests.post(OPTIMAILAPI + "/summarize", json={"content": body, "email": sender_email})
            print("**LOG**:", r.status_code, r.text)
            r = requests.post(OPTIMAILAPI + "/emphasize", json={"content": body, "email": sender_email})
            print("**LOG**:", r.status_code, r.text)
            mail.store(num, "+FLAGS", "\\Deleted")
        mail.expunge()
        time.sleep(1)
