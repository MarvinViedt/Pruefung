package main

const webseite = `
	<html>
		<center>
		<head>
			<div class="header-container">
			<center>
			<link rel="stylesheet" type="text/css" href="/static/css/Webseite.css">
			<title>CAMT053-Webservice</title>
			<h1>CAMT053 - Webservice</h1>
			</center>
			</div>
		</head>
		<body>
			<br></br>
			<div class="container">
			<form class="text-center" action="/CAMT053" method="post">
				<fieldset>
					<legend>Bitte die XML Datei angeben:</legend>
					<label for="XMLFile">XML Datei:</label>
						<input id="XMLFile" name="XMLFile" type="text" size="60" placeholder="Dateiname">
				</fieldset>
				<br></br>
				<button type="submit" name="Einlesen" formaction="/Camt053/Kontoauszugs-View">Kontoauszug anzeigen</button>
				<button type="submit" name="Einlesen" formaction="/Camt053/XML-View">XML anzeigen</button>
			</form>
			</div>
		</body>
		</center>
	</html>
`

















