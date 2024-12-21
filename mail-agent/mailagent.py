import imaplib
import email
import requests
from email.header import decode_header
import time
from getpass import getpass
import traceback

IMAP_SERVER = "mail.abdtyx.click"
IMAP_PORT = 993
EMAIL = "admin@abdtyx.click"
PASSWORD = getpass()
OPTIMAILAPI = "https://optimail.abdtyx.click/api"

with imaplib.IMAP4_SSL(IMAP_SERVER, IMAP_PORT) as mail:
    mail.login(EMAIL, PASSWORD)
    while 1:
        try:
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
                            payload = part.get_payload(decode=True)
                            if isinstance(payload, bytes) and payload:
                                charset = part.get_content_charset() or "utf-8"
                                try:
                                    body += payload.decode(charset)
                                except (UnicodeDecodeError, LookupError):
                                    body += payload.decode("iso-8859-1", errors="replace")

                else:
                    payload = msg.get_payload(decode=True)
                    if isinstance(payload, bytes) and payload:
                        charset = msg.get_content_charset() or "utf-8"
                        try:
                            body += payload.decode(charset)
                        except (UnicodeDecodeError, LookupError):
                            body += payload.decode("iso-8859-1", errors="replace")

                # print("Body:", body)

                # call Optimail api
                r = requests.post(OPTIMAILAPI + "/summarize", json={"content": body, "email": sender_email})
                print("**LOG**:", r.status_code, r.text)
                r = requests.post(OPTIMAILAPI + "/emphasize", json={"content": body, "email": sender_email})
                print("**LOG**:", r.status_code, r.text)
                mail.store(num, "+FLAGS", "\\Deleted")
            mail.expunge()
            time.sleep(1)
        except Exception as e:
            print(e)
            traceback.print_exc()
            time.sleep(1)
            continue
