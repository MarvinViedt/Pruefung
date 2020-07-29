package main

const KontoauszugsView = `
	<html>
		<center>
		<head>
			<div class="header-container">
			<link rel="stylesheet" type="text/css" href="/static/css/Kontoauszugs-View.css">
			<center>
			<title>CAMT053 - Webservice</title>
			<h1>CAMT053 - Webservice</h1>
			</center>
			</div>
		</head>
		</center>
		<body>
			<h4>Kontoinformationen</h4>
			<div class="kontoinfo-container">
				<div class="box2">
					<h4>Konto-Nr.</h4>
					<p>{{.Kntnr}}</p>
					<h4>BLZ</h4>
					<p>{{.BLZ}}</p>
				</div>
				<div class="box3">
					<h4>IBAN</h4>
					<p>{{.IBAN}}</p>
					<h4>BIC</h4>
					<p>{{.BIC}}</p>	
				</div>
				<div class="box4">
					<h3>Kontostand am {{.CreateDate}}</h3>
				</div>
				<div class="box4">
					<h3>{{.Kntst}} EUR</h3>
				</div>
				<div style="clear:both;"></div>
			</div>
			<h4>Auswahlkriterien</h4>
			<div class="kriterien-container">
				<p>{{.CreateDate}}</p>
			</div>
			<h4>Umsätze im gewählten Zeitraum</h4>
			<div class="table-container">
			<table>
			<tr>
				<th>Buchung</th>
				<th>Valuta</th>
				<th>Verwendungszweck</th>
				<th>Betrag</th>
			</tr>
			<tr>
				<td>{{.buchdt}}</td>
				<td>{{.valdt}}</td>
				<td>
				<b>{{.zweck}}</b>
				<br>
				{{.zwecknm}}
				<br>
				{{.info}}
				<br>
				KUNDENREFERENZ {{.kndref}}
				<br>
				MANDATSREFERENZ {{.manref}}
				</td>
				<td>{{.ums}} EUR</td>
			</tr>
			</table>
			</div>
			<br></br>
			<center>
			<form class="text-center" action="/Camt053/Kontoauszugs-View" method="post">
				<button type="submit" name="Zurück" formaction="/">Zurück</button>
			</form>
			</center>
		</body>
	</html>
`
