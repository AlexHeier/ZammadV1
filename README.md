# BedKomBedpressInvSender

Automatiserer det å sende standard invitasjonens mailen til Login - BedKom ved å ta over tastaturet og musen til maskinen og fyller inn innformasjonen utifra en .csv fil.

## Before use

Den er laget for opera med vanlig innstilliger på Windows 11.
Skjerm oppløsning 1440p!

Sørg for at du kan kjøre golang kode.

Sørg for at du har gått inn i main.go og forandrett pathen til csv filen.
Sørg for at csv filen følger riktig format (se lengere ned)
Sørg for at det står riktig eier i EmailOwner

Nå er det eneste som mangler er å gjøre om ToSend til true. Den kommer i test modus.

## CSV file upset

Dette er "header"en til .csv filen.

```plaintext
Bedrift, Emails, CC
```

Bedrift er bedrifts navnet.
Email kan KUNN være en av og er mottakeren.
CC er alle andre som skal ha mailen. Her kan det være flere.